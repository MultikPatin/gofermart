package interfaces

import "context"

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
	Register(ctx context.Context) error
	Login(ctx context.Context) error
}
