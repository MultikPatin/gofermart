package schemas

//easyjson -all internal/schemas/balance.go

type Balance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

type Withdraw struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

type Withdrawal struct {
	Order     string  `json:"order"`
	Sum       float32 `json:"sum"`
	Processed string  `json:"processed_at"`
}
