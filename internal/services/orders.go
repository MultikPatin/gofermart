package services

import (
	"context"
	"main/internal/dtos"
	"main/internal/interfaces"
	"time"
)

type OrdersService struct {
	repo interfaces.OrdersRepository
}

func NewOrdersService(r interfaces.OrdersRepository) *OrdersService {
	return &OrdersService{
		repo: r,
	}
}

func (s *OrdersService) Add(ctx context.Context, OrderID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.repo.Add(ctx)
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
