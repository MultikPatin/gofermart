package app

import (
	"github.com/go-chi/chi/v5"
	"main/internal/adapters/database/postgres"
	"main/internal/adapters/database/postgres/repositories"
	"main/internal/config"
	"main/internal/interfaces"
	"main/internal/services"
)

type App struct {
	Addr     string
	Router   *chi.Mux
	Services *Services
}

func (a *App) Close() error {
	err := a.Services.Close()
	if err != nil {
		return err
	}
	return nil
}

type Handlers struct {
	users    interfaces.UsersHandler
	orders   interfaces.OrdersHandler
	balances interfaces.BalancesHandler
}

type Services struct {
	users      interfaces.UsersService
	orders     interfaces.OrdersService
	balances   interfaces.BalancesService
	Repository *Repositories
}

func (s *Services) Close() error {
	err := s.Repository.Close()
	if err != nil {
		return err
	}
	return nil
}

type Repositories struct {
	users    interfaces.UsersRepository
	orders   interfaces.OrdersRepository
	balances interfaces.BalancesRepository
	Database interfaces.Database
}

func (s *Repositories) Close() error {
	err := s.Database.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewApp(c *config.Config) (*App, error) {
	s, err := NewServices(c)
	if err != nil {
		return nil, err
	}
	h := NewHandlers(s)
	r := NewRouters(h)

	return &App{
		Addr:     c.Addr,
		Router:   r,
		Services: s,
	}, nil
}

func NewHandlers(s *Services) *Handlers {
	return &Handlers{
		orders:   NewOrdersHandler(s.orders),
		balances: NewBalancesHandler(s.balances),
		users:    NewUsersHandler(s.users),
	}
}

func NewServices(c *config.Config) (*Services, error) {
	r, err := NewRepositories(c)
	if err != nil {
		return nil, err
	}
	return &Services{
		orders:     services.NewOrdersService(r.orders),
		balances:   services.NewBalancesService(r.balances),
		users:      services.NewUsersService(r.users),
		Repository: r,
	}, nil
}

func NewRepositories(c *config.Config) (*Repositories, error) {
	db, err := postgres.NewDatabase(c.PostgresDNS)
	if err != nil {
		return nil, err
	}
	return &Repositories{
		orders:   repositories.NewOrdersRepository(db),
		users:    repositories.NewUsersRepository(db),
		balances: repositories.NewBalancesRepository(db),
		Database: db,
	}, nil
}
