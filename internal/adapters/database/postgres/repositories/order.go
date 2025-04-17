package repositories

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/adapters/database/postgres"
	"main/internal/constants"
	"main/internal/services"
)

func NewOrdersRepository(db *postgres.Database) *OrdersRepository {
	return &OrdersRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type OrdersRepository struct {
	db     *postgres.Database
	logger *zap.SugaredLogger
}

func (r *OrdersRepository) Add(ctx context.Context, OrderNumber string) (int64, error) {
	userIDContext := ctx.Value(constants.UserIDKey).(int64)
	orderExist := true

	query := `SELECT user_id, order_id FROM orders WHERE order_id = $1`

	var userID int64
	var orderID int64
	row := r.db.Connection.QueryRowContext(ctx, query, OrderNumber)
	err := row.Scan(&userID, &orderID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			orderExist = false
		default:
			return -1, err
		}
	}

	if orderExist {
		if userID != userIDContext {
			return -1, services.ErrOrderAlreadyLoadedByAnotherUser
		} else {
			return -1, services.ErrOrderAlreadyExists
		}
	}

	query = `INSERT INTO orders (user_id, order_id) VALUES ($1, $2) RETURNING id`

	err = r.db.Connection.QueryRowContext(ctx, query, userIDContext, OrderNumber).Scan(&orderID)
	if err != nil {
		return -1, err
	}
	return orderID, err
}

func (r *OrdersRepository) GetAll(ctx context.Context) error {

	return nil
}
