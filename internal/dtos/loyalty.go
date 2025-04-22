package dtos

import "main/internal/enums"

type Deposit struct {
	OrderNumber string
	Amount      float32
}

type LoyaltyCalculation struct {
	Status  enums.LoyaltyStatusEnum
	Accrual float32
}

type UpdateOrderStatus struct {
	Number  string
	ID      int64
	Status  enums.LoyaltyStatusEnum
	Accrual float32
}
