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

func NewBalancesRepository(db *postgres.Database) *BalancesRepository {
	return &BalancesRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type BalancesRepository struct {
	db     *postgres.Database
	logger *zap.SugaredLogger
}

func (r *BalancesRepository) Get(ctx context.Context) (*dtos.Balance, error) {
	userID := ctx.Value(constants.UserIDKey).(int64)

	query := `
	SELECT
		SUM(CASE WHEN action = $1 THEN amount WHEN action = $2 THEN -amount END) AS current,
    	SUM(CASE WHEN action = $2 THEN amount ELSE 0 END) AS withdrawn
	FROM balances
	WHERE user_id = $3;`

	var result dtos.Balance
	row := r.db.Connection.QueryRowContext(ctx, query, enums.BalanceDeposit.String(), enums.BalanceWithdrawal.String(), userID)
	err := row.Scan(&result.Current, &result.Withdraw)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BalancesRepository) Withdraw(ctx context.Context, withdrawal *dtos.Withdraw) (int64, error) {
	userID := ctx.Value(constants.UserIDKey).(int64)
	action := enums.BalanceWithdrawal.String()
	var ID int64

	orderExist := true

	query := `
	SELECT user_id
	FROM orders
	WHERE order_id = $1 AND user_id = $2;`

	row := r.db.Connection.QueryRowContext(ctx, query, withdrawal.Order, userID)
	err := row.Scan(&ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			orderExist = false
		default:
			return -1, err
		}
	}

	if !orderExist {
		return -1, services.ErrOrderIDNotValid
	}

	query = `
	INSERT INTO balances (user_id, order_id, action, amount)
	VALUES ($1, $2, $3, $4) RETURNING id;`

	err = r.db.Connection.QueryRowContext(ctx, query, userID, withdrawal.Order, action, withdrawal.Sum).Scan(&ID)
	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (r *BalancesRepository) Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error) {
	userID := ctx.Value(constants.UserIDKey).(int64)
	action := enums.BalanceWithdrawal.String()

	query := `
    SELECT order_id, amount, processed_at
    FROM balances
    WHERE action = $1 and user_id = $2
    ORDER BY processed_at DESC;`

	rows, err := r.db.Connection.QueryContext(ctx, query, action, userID)
	if err != nil {
		return nil, err
	}

	var processedAt time.Time
	var withdrawals []*dtos.Withdrawal

	for rows.Next() {
		var w dtos.Withdrawal
		err := rows.Scan(&w.Order, &w.Sum, &processedAt)
		if err != nil {
			return nil, err
		}
		w.Processed = processedAt.Format(time.RFC3339)
		withdrawals = append(withdrawals, &w)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return withdrawals, nil
}

func (r *BalancesRepository) BatchAdd(ctx context.Context, items []*dtos.Deposit) ([]int64, error) {
	userID := ctx.Value(constants.UserIDKey).(int64)
	action := enums.BalanceDeposit.String()

	tx, err := r.db.Connection.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO balances (user_id, order_id, action, amount)
	VALUES ($1, $2, $3, $4) RETURNING id;`

	var results []int64

	for _, item := range items {
		var ID int64
		err := r.db.Connection.QueryRowContext(ctx, query, userID, item.OrderNumber, action, item.Amount).Scan(&ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		results = append(results, ID)
	}

	tx.Commit()
	return results, nil
}
