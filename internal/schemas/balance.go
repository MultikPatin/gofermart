package schemas

//easyjson -all internal/schemas/balance.go

type Balance struct {
	Current  float32 `json:"current"`
	Withdraw float32 `json:"withdraw"`
}

type Withdraw struct {
	Order int     `json:"order"`
	Sum   float32 `json:"sum"`
}

type Withdrawal struct {
	Order     int     `json:"order"`
	Sum       float32 `json:"sum"`
	Processed string  `json:"processed_at"`
}
