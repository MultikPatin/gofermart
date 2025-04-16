package interfaces

import "main/internal/dtos"

type Database interface {
	Close() error
	Ping() error
}

type LoyaltyCalculation interface {
	GetByOrderID(orderID string) (*dtos.LoyaltyCalculation, error)
}
