package interfaces

type BalanceService interface {
	Get()
	Withdraw()
}

type OrdersService interface {
	Add()
	GetAll()
}

type UsersService interface {
	Register()
	Login()
	Withdrawals()
}
