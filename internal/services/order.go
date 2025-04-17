package services

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/dtos"
	"main/internal/interfaces"
	"time"
)

var (
	ErrOrderAlreadyExists              = errors.New("order already exists")
	ErrOrderAlreadyLoadedByAnotherUser = errors.New("order already loaded by another user")
	ErrOrderIDNotValid                 = errors.New("order id is not valid")
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

	err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
