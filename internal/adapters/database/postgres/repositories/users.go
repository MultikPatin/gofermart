package repositories

import (
	"context"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/interfaces"
)

func NewUsersRepository(db interfaces.Database) *UsersRepository {
	return &UsersRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type UsersRepository struct {
	db     interfaces.Database
	logger *zap.SugaredLogger
}

func (s *UsersRepository) Login(ctx context.Context) error {

	return nil
}

func (s *UsersRepository) Register(ctx context.Context) error {

	return nil
}
