package repositories

import (
	"context"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/interfaces"
)

func NewBalancesRepository(db interfaces.Database) *BalancesRepository {
	return &BalancesRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type BalancesRepository struct {
	db     interfaces.Database
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
