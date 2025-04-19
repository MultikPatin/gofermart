package services

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/dtos"
	"main/internal/interfaces"
	"sync"
	"time"
)

var (
	ErrOrderAlreadyExists              = errors.New("order already exists")
	ErrOrderAlreadyLoadedByAnotherUser = errors.New("order already loaded by another user")
	ErrOrderIDNotValid                 = errors.New("order id is not valid")
	ErrTooManyRequests                 = errors.New("too many requests to the client")
)

func NewOrdersService(r interfaces.OrdersRepository, lc interfaces.LoyaltyCalculation) *OrdersService {
	return &OrdersService{
		repo:   r,
		logger: adapters.GetLogger(),
		lc:     lc,
	}
}

type OrdersService struct {
	repo   interfaces.OrdersRepository
	logger *zap.SugaredLogger
	lc     interfaces.LoyaltyCalculation
}

func (s *OrdersService) Add(ctx context.Context, OrderID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//TODO Номер заказа может быть проверен на корректность ввода с помощью алгоритма Луна.

	_, err := s.repo.Add(ctx, OrderID)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrdersService) GetAll(ctx context.Context) ([]*dtos.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//Доступные статусы обработки расчётов:
	//NEW — заказ загружен в систему, но не попал в обработку;
	//PROCESSING — вознаграждение за заказ рассчитывается;
	//INVALID — система расчёта вознаграждений отказала в расчёте;
	//PROCESSED — данные по заказу проверены и информация о расчёте успешно получена.

	orders, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []*dtos.Order

	ordersChan := ordersGenerator(ctx, orders)
	errChan := make(chan error, len(orders))
	resultChan := make(chan *dtos.Order, len(orders))

	var wg sync.WaitGroup
	for orderNumber := range ordersChan {
		wg.Add(1)
		data := orderNumber
		go func(orderDB *dtos.OrderDB) {
			defer wg.Done()
			loyalty, err := s.lc.GetByOrderID(ctx, orderDB.Number)
			if err != nil {
				errChan <- err
			} else {
				resultChan <- &dtos.Order{
					OrderDB:     *orderDB,
					OrderStatus: loyalty.OrderStatus,
				}
			}

		}(data)
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(resultChan)
	}()

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("get loyalty by order ids: %v", errs)
	}

	for item := range resultChan {
		results = append(results, item)
	}

	return results, nil
}

func ordersGenerator(ctx context.Context, orders []*dtos.OrderDB) chan *dtos.OrderDB {
	inputCh := make(chan *dtos.OrderDB, len(orders))

	go func() {
		defer close(inputCh)

		for _, order := range orders {
			select {
			case <-ctx.Done():
				return
			case inputCh <- order:
			}
		}
	}()
	return inputCh
}
