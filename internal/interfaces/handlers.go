package interfaces

import "net/http"

type BalancesHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
}

type OrdersHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type UsersHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Withdrawals(w http.ResponseWriter, r *http.Request)
}
