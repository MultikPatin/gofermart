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
	Accrual  float32
}
