package services

import (
	"context"
	"main/internal/interfaces"
	"time"
)

type UsersService struct {
	repo interfaces.UsersRepository
}

func NewUsersService(r interfaces.UsersRepository) *UsersService {
	return &UsersService{
		repo: r,
	}
}

func (s *UsersService) Login(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.repo.Login(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) Register(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := s.repo.Register(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) Withdrawals(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := s.repo.Withdrawals(ctx)
	if err != nil {
		return err
	}

	return nil
}
