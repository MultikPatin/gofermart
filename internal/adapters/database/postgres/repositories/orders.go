package repositories

import (
	"context"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/interfaces"
)

func NewOrdersRepository(db interfaces.Database) *OrdersRepository {
	return &OrdersRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type OrdersRepository struct {
	db     interfaces.Database
	logger *zap.SugaredLogger
}

func (s *OrdersRepository) Add(ctx context.Context) error {

	return nil
}

func (s *OrdersRepository) GetAll(ctx context.Context) error {

	return nil
}
