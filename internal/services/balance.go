package services

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/dtos"
	"main/internal/interfaces"
	"time"
)

var (
	ErrPaymentRequired = errors.New("no more balance to withdraw")
)

func NewBalancesService(r interfaces.BalancesRepository, loyaltyService interfaces.LoyaltyService) *BalancesService {
	return &BalancesService{
		repo:   r,
		logger: adapters.GetLogger(),
		ls:     loyaltyService,
	}
}

type BalancesService struct {
	repo   interfaces.BalancesRepository
	logger *zap.SugaredLogger
	ls     interfaces.LoyaltyService
}

func (s *BalancesService) Get(ctx context.Context) (*dtos.Balance, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.ls.Update(ctx)
	if err != nil {
		return nil, err
	}

	balance, err := s.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (s *BalancesService) Withdraw(ctx context.Context, withdrawal *dtos.Withdraw) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	balance, err := s.Get(ctx)
	if err != nil {
		return err
	}

	if balance.Current < withdrawal.Sum {
		return ErrPaymentRequired
	}

	_, err = s.repo.Withdraw(ctx, withdrawal)
	if err != nil {
		return err
	}

	return nil
}

func (s *BalancesService) Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	results, err := s.repo.Withdrawals(ctx)
	if err != nil {
		return nil, err
	}

	return results, nil
}
