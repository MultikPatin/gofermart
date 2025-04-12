package repositories

import (
	"context"
	"main/internal/interfaces"
)

type UsersRepository struct {
	db interfaces.Database
}

func NewUsersRepository(db interfaces.Database) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (s *UsersRepository) Login(ctx context.Context) error {

	return nil
}

func (s *UsersRepository) Register(ctx context.Context) error {

	return nil
}

func (s *UsersRepository) Withdrawals(ctx context.Context) error {

	return nil
}
