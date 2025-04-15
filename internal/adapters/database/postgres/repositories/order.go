package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/adapters/database/postgres"
)

func NewOrdersRepository(db *postgres.Database) *OrdersRepository {
	return &OrdersRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type OrdersRepository struct {
	db     *postgres.Database
	logger *zap.SugaredLogger
}

func (s *OrdersRepository) Add(ctx context.Context) error {

	return nil
}

func (s *OrdersRepository) GetAll(ctx context.Context) error {

	return nil
}
