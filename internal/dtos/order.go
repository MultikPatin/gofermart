package dtos

type Order struct {
	Number   string
	Status   string
	Accrual  float32
	Uploaded string
}

type OrderAccrual struct {
	Order   int
	Status  string
	Accrual float32
}
