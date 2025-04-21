package interfaces

import (
	"context"
	"main/internal/dtos"
)

type BalancesService interface {
	Get(ctx context.Context) (*dtos.Balance, error)
	Withdraw(ctx context.Context, withdrawal *dtos.Withdraw) error
	Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error)
}

type OrdersService interface {
	Add(ctx context.Context, OrderID string) error
	GetAll(ctx context.Context) ([]*dtos.OrderDB, error)
}

type UsersService interface {
	Register(ctx context.Context, credentials *dtos.AuthCredentials) (int64, error)
	Login(ctx context.Context, credentials *dtos.AuthCredentials) (int64, error)
}

type LoyaltyService interface {
	Update(ctx context.Context) error
}
