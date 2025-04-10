package interfaces

import "net/http"

type BalanceHandlers interface {
	Get(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
}

type OrdersHandlers interface {
	Add(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type UsersHandlers interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Withdrawals(w http.ResponseWriter, r *http.Request)
}
