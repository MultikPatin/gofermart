package repositories

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/adapters/database/postgres"
	"main/internal/constants"
	"main/internal/dtos"
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
		SUM(CASE WHEN action = 'deposit' THEN amount WHEN action = 'withdrawal' THEN -amount END) AS current,
    	SUM(CASE WHEN action = 'withdrawal' THEN amount ELSE 0 END) AS withdrawn
	FROM balances
	WHERE user_id = $1;`

	result := new(dtos.Balance)
	row := r.db.Connection.QueryRowContext(ctx, query, userID)
	err := row.Scan(&result.Current, &result.Withdraw)
	if err == nil {
		return nil, err
	}

	return result, nil
}

func (r *BalancesRepository) Withdraw(ctx context.Context, withdrawal *dtos.Withdraw) error {
	// если заказ не найден в таблице orders
	//err := services.ErrOrderIDNotValid
	return nil
}

func (r *BalancesRepository) Withdrawals(ctx context.Context) ([]*dtos.Withdrawal, error) {
	userID := ctx.Value(constants.UserIDKey).(int64)

	query := `
    SELECT order_id, amount, processed_at
    FROM balances
    WHERE action = 'withdrawal' and user_id = $1
    ORDER BY processed_at DESC;`

	rows, err := r.db.Connection.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var processedAt time.Time
	var withdrawals []*dtos.Withdrawal

	for rows.Next() {
		w := new(dtos.Withdrawal)
		err := rows.Scan(&w.Order, &w.Sum, &processedAt)
		if err != nil {
			return nil, err
		}
		w.Processed = processedAt.Format(time.RFC3339)
		withdrawals = append(withdrawals, w)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return withdrawals, nil
}
