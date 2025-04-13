package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/dtos"
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

func (s *UsersRepository) GetByLogin(ctx context.Context, login string) (dtos.User, error) {

	return dtos.User{}, nil
}

func (s *UsersRepository) Add(ctx context.Context, credentials dtos.AuthCredentials) (int64, error) {

	return 0, nil
}
