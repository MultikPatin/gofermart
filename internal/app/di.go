package app

import (
	"github.com/go-chi/chi/v5"
	"main/internal/adapters"
	"main/internal/adapters/db/psql"
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
	users    interfaces.UsersHandlers
	orders   interfaces.OrdersHandlers
	balances interfaces.BalanceHandlers
}

type Services struct {
	users      interfaces.UsersService
	orders     interfaces.OrdersService
	balances   interfaces.BalanceService
	Repository *Repository
}

func (s *Services) Close() error {
	err := s.Repository.Close()
	if err != nil {
		return err
	}
	return nil
}

type Repository struct {
	users    interfaces.UsersRepository
	orders   interfaces.OrdersRepository
	balances interfaces.BalanceRepository
	Database interfaces.DBConnection
}

func (s *Repository) Close() error {
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
		orders:   NewOrdersHandlers(s.orders),
		balances: NewBalancesHandlers(s.balances),
		users:    NewUsersHandlers(s.users),
	}
}

func NewServices(c *config.Config) (*Services, error) {
	repository, err := NewRepository(c)
	if err != nil {
		return nil, err
	}
	return &Services{
		orders:     services.NewOrdersService(c, repository.orders),
		balances:   services.NewBalancesService(repository.balances),
		users:      services.NewUserService(repository.users),
		Repository: repository,
	}, nil
}

func NewRepository(c *config.Config) (*Repository, error) {
	var repository *Repository

	logger := adapters.GetLogger()

	db, err := psql.NewPostgresDB(c.PostgresDNS, logger)
	if err != nil {
		return nil, err
	}
	logger.Info("Create PostgresDB Connection")
	repository = NewPostgresRepository(db)
	return repository, nil
}

func NewPostgresRepository(db *psql.PostgresDB) *Repository {
	return &Repository{
		orders:   psql.NewOrdersRepository(db),
		users:    psql.NewUsersRepository(db),
		balances: psql.NewBalancesRepository(db),
		Database: db,
	}
}
