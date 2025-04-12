package services

import (
	"context"
	"main/internal/dtos"
	"main/internal/interfaces"
	"time"
)

type BalancesService struct {
	repo interfaces.BalancesRepository
}

func NewBalancesService(r interfaces.BalancesRepository) *BalancesService {
	return &BalancesService{
		repo: r,
	}
}

func (s *BalancesService) Get(ctx context.Context) (dtos.Balance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *BalancesService) Withdraw(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := s.repo.Withdraw(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *BalancesService) Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := s.repo.Withdrawals(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
