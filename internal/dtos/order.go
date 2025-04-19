package dtos

type OrderBase struct {
	Number string
}
type OrderStatus struct {
	Status  string
	Accrual float32
}

type OrderBD struct {
	ID int64
	OrderBase
	Uploaded string
}

type LoyaltyCalculation struct {
	OrderBase
	OrderStatus
}

type Order struct {
	OrderBD
	OrderStatus
}
