package dtos

type Order struct {
	Number   string
	Status   string
	Accrual  float32
	Uploaded string
}

type LoyaltyCalculation struct {
	Number  string
	Status  string
	Accrual float32
}

type OrderAccrual struct {
	Order   int
	Status  string
	Accrual float32
}
