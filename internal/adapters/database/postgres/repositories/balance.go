package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/adapters/database/postgres"
)

func NewBalancesRepository(db *postgres.Database) *BalancesRepository {
	return &BalancesRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type BalancesRepository struct {
	db     *postgres.Database
	logger *zap.SugaredLogger
}

func (s *BalancesRepository) Get(ctx context.Context) error {

	return nil
}

func (s *BalancesRepository) Withdraw(ctx context.Context) error {

	return nil
}

func (s *BalancesRepository) Withdrawals(ctx context.Context) error {

	return nil
}
