package services

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/constants"
	"main/internal/dtos"
	"main/internal/enums"
	"main/internal/interfaces"
	"sync"
)

func NewLoyaltyService(ro interfaces.OrdersRepository, rb interfaces.BalancesRepository, lc interfaces.LoyaltyCalculation) *LoyaltyService {
	return &LoyaltyService{
		ro:     ro,
		rb:     rb,
		logger: adapters.GetLogger(),
		lc:     lc,
	}
}

type LoyaltyService struct {
	ro     interfaces.OrdersRepository
	rb     interfaces.BalancesRepository
	logger *zap.SugaredLogger
	lc     interfaces.LoyaltyCalculation
}

func (s *LoyaltyService) AddBalances(ctx context.Context, balances [][]*dtos.Deposit) error {
	inputChan := depositGenerator(ctx, balances)
	errChan := make(chan error, len(balances))

	var wg sync.WaitGroup
	for item := range inputChan {
		wg.Add(1)
		data := item
		go func(balanceAdd []*dtos.Deposit) {
			defer wg.Done()
			_, err := s.rb.BatchAdd(ctx, balanceAdd)
			if err != nil {
				errChan <- fmt.Errorf("error updating balances: %w", err)
			}
		}(data)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		s.logger.Infow(
			"Update balances",
			"errors", err.Error(),
		)
	}

	return nil
}

func (s *LoyaltyService) UpdateOrders(ctx context.Context, orders []*dtos.UpdateOrderStatus) error {

	var batches [][]*dtos.UpdateOrderStatus

	for i := 0; i < len(orders); i += constants.LoyaltyBatchSize {
		end := i + constants.LoyaltyBatchSize
		if end > len(orders) {
			end = len(orders)
		}
		batch := orders[i:end]
		batches = append(batches, batch)
	}

	var results [][]*dtos.Deposit

	inputChan := ordersBatchGenerator(ctx, batches)
	errChan := make(chan error, len(batches))
	resultChan := make(chan []*dtos.Deposit, len(batches))

	var wg sync.WaitGroup
	for item := range inputChan {
		wg.Add(1)
		data := item
		go func(orderUpdate []*dtos.UpdateOrderStatus) {
			defer wg.Done()
			s.logger.Infow(
				"BatchUpdate",
				"orderUpdate", orderUpdate,
			)
			err := s.ro.BatchUpdate(ctx, orderUpdate)
			if err != nil {
				errChan <- fmt.Errorf("error updating orders: %w", err)
			} else {
				results := make([]*dtos.Deposit, len(orderUpdate))
				for _, order := range orderUpdate {
					if order.Status == enums.LoyaltyProcessed {
						result := &dtos.Deposit{
							OrderNumber: order.Number,
							Amount:      order.Accrual,
						}
						results = append(results, result)
					}
				}
				if len(results) > 0 {
					resultChan <- results
				}
			}
		}(data)
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(resultChan)
	}()

	for err := range errChan {
		s.logger.Infow(
			"Update orders",
			"errors", err.Error(),
		)
	}

	for item := range resultChan {
		results = append(results, item)
	}

	if len(results) > 0 {
		err := s.AddBalances(ctx, results)
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *LoyaltyService) Update(ctx context.Context) error {

	statuses := []enums.OrderStatusEnum{
		enums.OrderCreated,
		enums.OrderProcessing,
	}

	orders, err := s.ro.GetAll(ctx, statuses)
	if err != nil {
		return err
	}

	results := make([]*dtos.UpdateOrderStatus, len(orders))

	inputChan := ordersGenerator(ctx, orders)
	errChan := make(chan error, len(orders))
	resultChan := make(chan *dtos.UpdateOrderStatus, len(orders))

	var wg sync.WaitGroup
	for item := range inputChan {
		wg.Add(1)
		data := item
		go func(orderDB *dtos.OrderDB) {
			defer wg.Done()
			loyalty, err := s.lc.GetByOrderID(ctx, orderDB.Number)
			if err != nil {
				errChan <- err
			} else {
				resultChan <- &dtos.UpdateOrderStatus{
					ID:      orderDB.ID,
					Status:  loyalty.Status,
					Accrual: loyalty.Accrual,
					Number:  orderDB.Number,
				}
			}

		}(data)
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(resultChan)
	}()

	for err := range errChan {
		s.logger.Infow(
			"Update loyalty",
			"errors", err.Error(),
		)
	}

	for item := range resultChan {
		results = append(results, item)
	}

	err = s.UpdateOrders(ctx, results)
	if err != nil {
		return err
	}

	return nil

}

func depositGenerator(ctx context.Context, items [][]*dtos.Deposit) chan []*dtos.Deposit {
	inputCh := make(chan []*dtos.Deposit, len(items))

	go func() {
		defer close(inputCh)

		for _, item := range items {
			select {
			case <-ctx.Done():
				return
			case inputCh <- item:
			}
		}
	}()
	return inputCh
}

func ordersGenerator(ctx context.Context, items []*dtos.OrderDB) chan *dtos.OrderDB {
	inputCh := make(chan *dtos.OrderDB, len(items))

	go func() {
		defer close(inputCh)

		for _, item := range items {
			select {
			case <-ctx.Done():
				return
			case inputCh <- item:
			}
		}
	}()
	return inputCh
}

func ordersBatchGenerator(ctx context.Context, items [][]*dtos.UpdateOrderStatus) chan []*dtos.UpdateOrderStatus {
	inputCh := make(chan []*dtos.UpdateOrderStatus, len(items))

	go func() {
		defer close(inputCh)

		for _, item := range items {
			select {
			case <-ctx.Done():
				return
			case inputCh <- item:
			}
		}
	}()
	return inputCh
}
