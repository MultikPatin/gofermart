package interfaces

import (
	"context"
	"main/internal/dtos"
)

type BalancesRepository interface {
	Get(ctx context.Context) (*dtos.Balance, error)
	Withdraw(ctx context.Context, withdrawal *dtos.Withdraw) error
	Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error)
}

type OrdersRepository interface {
	Add(ctx context.Context, OrderID string) (int64, error)
	GetAll(ctx context.Context) ([]*dtos.OrderDB, error)
}

type UsersRepository interface {
	Add(ctx context.Context, credentials *dtos.AuthCredentials) (int64, error)
	GetByLogin(ctx context.Context, login string) (*dtos.User, error)
}
