package interfaces

type BalanceRepository interface {
	Get()
	Withdraw()
}

type OrdersRepository interface {
	Add()
	GetAll()
}

type UsersRepository interface {
	Register()
	Login()
	Withdrawals()
}
