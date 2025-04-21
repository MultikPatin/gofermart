package services

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/dtos"
	"main/internal/enums"
	"main/internal/interfaces"
	"strconv"
	"time"
)

var (
	ErrOrderAlreadyExists              = errors.New("order already exists")
	ErrOrderAlreadyLoadedByAnotherUser = errors.New("order already loaded by another user")
	ErrOrderIDNotValid                 = errors.New("order id is not valid")
	ErrTooManyRequests                 = errors.New("too many requests to the client")
	ErrUnknownStatus                   = errors.New("unknown order status")
)

func NewOrdersService(r interfaces.OrdersRepository, loyaltyService interfaces.LoyaltyService) *OrdersService {
	return &OrdersService{
		repo:   r,
		logger: adapters.GetLogger(),
		ls:     loyaltyService,
	}
}

type OrdersService struct {
	repo   interfaces.OrdersRepository
	logger *zap.SugaredLogger
	ls     interfaces.LoyaltyService
}

func (s *OrdersService) Add(ctx context.Context, OrderID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if !IsValidLuhn(OrderID) {
		return ErrOrderIDNotValid
	}

	orderCreate := &dtos.OrderCreate{
		Number: OrderID,
		Status: enums.OrderCreated,
	}

	_, err := s.repo.Add(ctx, orderCreate)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrdersService) GetAll(ctx context.Context) ([]*dtos.OrderDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.ls.Update(ctx)
	if err != nil {
		return nil, err
	}

	var statuses []enums.OrderStatusEnum

	orders, err := s.repo.GetAll(ctx, statuses)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func IsValidLuhn(number string) bool {
	sum := 0
	double := false

	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}
