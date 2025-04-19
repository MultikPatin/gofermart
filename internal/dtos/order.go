package dtos

type OrderBase struct {
	Number string
}
type OrderStatus struct {
	Status  string
	Accrual float32
}

type OrderDB struct {
	ID int64
	OrderBase
	Uploaded string
}

type LoyaltyCalculation struct {
	OrderBase
	OrderStatus
}

type Order struct {
	OrderDB
	OrderStatus
}
