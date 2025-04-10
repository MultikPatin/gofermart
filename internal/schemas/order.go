package schemas

type Order struct {
	Number   string  `json:"number"`
	Status   string  `json:"status"`
	Accrual  float32 `json:"accrual"`
	Uploaded string  `json:"uploaded_at"`
}

type OrderAccrual struct {
	Order   int     `json:"order"`
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual,omitempty"`
}
