package clients

import "main/internal/dtos"

type LoyaltyCalculation struct {
	accrualSystemAddr string
}

func NewLoyaltyCalculation(Addr string) *LoyaltyCalculation {
	return &LoyaltyCalculation{
		accrualSystemAddr: Addr,
	}
}

func (l *LoyaltyCalculation) GetByOrderID(orderID string) (*dtos.LoyaltyCalculation, error) {
	return nil, nil
}
