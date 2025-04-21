package interfaces

import (
	"context"
	"main/internal/dtos"
	"main/internal/enums"
)

type BalancesRepository interface {
	Get(ctx context.Context) (*dtos.Balance, error)
	Withdraw(ctx context.Context, withdrawal *dtos.Withdraw) (int64, error)
	Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error)
	BatchAdd(ctx context.Context, orders []*dtos.Deposit) ([]int64, error)
}

type OrdersRepository interface {
	Add(ctx context.Context, orderCreate *dtos.OrderCreate) (int64, error)
	GetAll(ctx context.Context, statuses []enums.OrderStatusEnum) ([]*dtos.OrderDB, error)
	BatchUpdate(ctx context.Context, orders []*dtos.UpdateOrderStatus) error
}

type UsersRepository interface {
	Add(ctx context.Context, credentials *dtos.AuthCredentials) (int64, error)
	GetByLogin(ctx context.Context, login string) (*dtos.User, error)
}
