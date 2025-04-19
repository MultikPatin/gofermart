package dtos

import "main/internal/enums"

type OrderBase struct {
	Number string
}

type OrderCreate struct {
	Status enums.OrderStatusEnum
	Number string
}
type OrderStatus struct {
	Status  enums.OrderStatusEnum
	Accrual float32
}

type OrderDB struct {
	ID int64
	OrderBase
	Status   enums.OrderStatusEnum
	Uploaded string
}

type LoyaltyCalculation struct {
	OrderBase
	OrderStatus
}

type Order struct {
	OrderDB
	Accrual float32
}
