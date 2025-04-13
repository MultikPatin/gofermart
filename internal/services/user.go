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
	ErrLoginAlreadyExists        = errors.New("login already exists")
	ErrAuthCredentialsIsNotValid = errors.New("login or password is not valid")
)

func NewUsersService(r interfaces.UsersRepository) *UsersService {
	return &UsersService{
		repo:   r,
		logger: adapters.GetLogger(),
	}
}

type UsersService struct {
	repo   interfaces.UsersRepository
	logger *zap.SugaredLogger
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
