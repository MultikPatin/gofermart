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
	"main/internal/dtos"
	"main/internal/enums"
	"main/internal/services"
	"time"
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

func (r *OrdersRepository) Add(ctx context.Context, orderCreate *dtos.OrderCreate) (int64, error) {
	userIDContext := ctx.Value(constants.UserIDKey).(int64)
	orderExist := true

	query := `
	SELECT user_id, order_id 
	FROM orders 
	WHERE order_id = $1;`

	var userID int64
	var orderID int64
	row := r.db.Connection.QueryRowContext(ctx, query, orderCreate.Number)
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

	query = `
	INSERT INTO orders (user_id, order_id, status) 
	VALUES ($1, $2, $3) 
	RETURNING id;`

	err = r.db.Connection.QueryRowContext(ctx, query, userIDContext, orderCreate.Number, orderCreate.Status.String()).Scan(&orderID)
	if err != nil {
		return -1, err
	}
	return orderID, err
}

func (r *OrdersRepository) GetAll(ctx context.Context) ([]*dtos.OrderDB, error) {
	userID := ctx.Value(constants.UserIDKey).(int64)

	query := `
	SELECT id, order_id, status, uploaded_at 
	FROM orders 
	WHERE user_id = $1;`

	rows, err := r.db.Connection.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var uploadedAt time.Time
	var status string
	var orders []*dtos.OrderDB

	for rows.Next() {
		w := new(dtos.OrderDB)
		err := rows.Scan(&w.ID, &w.Number, &status, &uploadedAt)
		if err != nil {
			return nil, err
		}
		w.Uploaded = uploadedAt.Format(time.RFC3339)

		var ok bool

		w.Status, ok = enums.OrdesStatusFromString(status)
		if !ok {
			r.logger.Infow(
				"Get order with unknown status",
				"status", status,
			)
		} else {
			orders = append(orders, w)
		}

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
