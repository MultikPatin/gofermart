package services

import (
	"context"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/dtos"
	"main/internal/interfaces"
	"time"
)

func NewBalancesService(r interfaces.BalancesRepository) *BalancesService {
	return &BalancesService{
		repo:   r,
		logger: adapters.GetLogger(),
	}
}

type BalancesService struct {
	repo   interfaces.BalancesRepository
	logger *zap.SugaredLogger
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

func (s *BalancesService) Withdraw(ctx context.Context, withdrawal dtos.Withdraw) error {
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
