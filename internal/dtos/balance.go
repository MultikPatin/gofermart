package dtos

type Balance struct {
	Current  float32
	Withdraw float32
}

type Withdraw struct {
	Order string
	Sum   float32
}

type Withdrawal struct {
	Order     string
	Sum       float32
	Processed string
}
