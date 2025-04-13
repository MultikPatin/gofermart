package interfaces

import (
	"context"
	"main/internal/dtos"
)

type BalancesRepository interface {
	Get(ctx context.Context) error
	Withdraw(ctx context.Context) error
	Withdrawals(ctx context.Context) error
}

type OrdersRepository interface {
	Add(ctx context.Context) error
	GetAll(ctx context.Context) error
}

type UsersRepository interface {
	Add(ctx context.Context, credentials dtos.AuthCredentials) (int64, error)
	GetByLogin(ctx context.Context, login string) (dtos.User, error)
}
