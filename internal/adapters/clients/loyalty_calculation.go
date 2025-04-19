package clients

import (
	"context"
	"main/internal/dtos"
)

type LoyaltyCalculation struct {
	accrualSystemAddr string
}

func NewLoyaltyCalculation(Addr string) *LoyaltyCalculation {
	return &LoyaltyCalculation{
		accrualSystemAddr: Addr,
	}
}

func (l *LoyaltyCalculation) GetByOrderID(ctx context.Context, orderID string) (*dtos.LoyaltyCalculation, error) {
	return nil, nil
}
