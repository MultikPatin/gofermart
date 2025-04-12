package repositories

import (
	"context"
	"main/internal/interfaces"
)

type OrdersRepository struct {
	db interfaces.Database
}

func NewOrdersRepository(db interfaces.Database) *OrdersRepository {
	return &OrdersRepository{
		db: db,
	}
}

func (s *OrdersRepository) Add(ctx context.Context) error {

	return nil
}

func (s *OrdersRepository) GetAll(ctx context.Context) error {

	return nil
}
