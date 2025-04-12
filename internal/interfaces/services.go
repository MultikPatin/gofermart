package interfaces

import (
	"context"
	"main/internal/dtos"
)

type BalancesService interface {
	Get(ctx context.Context) error
	Withdraw(ctx context.Context) error
}

type OrdersService interface {
	Add(ctx context.Context) error
	GetAll(ctx context.Context) error
}

type UsersService interface {
	Register(ctx context.Context, credentials dtos.AuthCredentials) error
	Login(ctx context.Context, credentials dtos.AuthCredentials) error
	Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error)
}
