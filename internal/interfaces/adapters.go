package interfaces

import (
	"context"
	"main/internal/dtos"
)

type Database interface {
	Close() error
	Ping() error
}

type LoyaltyCalculation interface {
	GetByOrderID(ctx context.Context, orderID string) (*dtos.LoyaltyCalculation, error)
}
