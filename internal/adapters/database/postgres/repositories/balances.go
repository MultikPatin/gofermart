package repositories

import (
	"context"
	"main/internal/interfaces"
)

type BalancesRepository struct {
	db interfaces.Database
}

func NewBalancesRepository(db interfaces.Database) *BalancesRepository {
	return &BalancesRepository{
		db: db,
	}
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
