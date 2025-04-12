package app

import (
	"github.com/go-chi/chi/v5"
	"main/internal/middleware"
)

func NewRouters(h *Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AccessLogger)
	r.Use(middleware.GZipper)
	//r.Use(middleware.Authentication)

	r.Route("/", func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Post("/register", h.users.Register)
				r.Post("/login", h.users.Login)
				r.Post("/orders", h.orders.Add)
				r.Get("/orders", h.orders.GetAll)
				r.Get("/withdrawals", h.balances.Withdrawals)
				r.Route("/balance", func(r chi.Router) {
					r.Get("/", h.balances.Get)
					r.Post("/withdraw", h.balances.Withdraw)
				})
			})
		})
	})

	return r
}
