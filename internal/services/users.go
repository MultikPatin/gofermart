package services

import (
	"context"
	"main/internal/dtos"
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

func (s *UsersService) Login(ctx context.Context, credentials dtos.AuthCredentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.repo.Login(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) Register(ctx context.Context, credentials dtos.AuthCredentials) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := s.repo.Register(ctx)
	if err != nil {
		return err
	}

	return nil
}
