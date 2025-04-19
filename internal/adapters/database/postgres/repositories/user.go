package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"main/internal/adapters"
	"main/internal/adapters/database/postgres"
	"main/internal/dtos"
	"main/internal/services"
)

func NewUsersRepository(db *postgres.Database) *UsersRepository {
	return &UsersRepository{
		db:     db,
		logger: adapters.GetLogger(),
	}
}

type UsersRepository struct {
	db     *postgres.Database
	logger *zap.SugaredLogger
}

func (r *UsersRepository) GetByLogin(ctx context.Context, login string) (*dtos.User, error) {
	query := `
	SELECT id, login, password 
	FROM users 
	WHERE login=$1 
	LIMIT 1;`

	user := new(dtos.User)
	row := r.db.Connection.QueryRowContext(ctx, query, login)
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, services.ErrAuthCredentialsIsNotValid
		default:
			return nil, err
		}
	}
	return user, nil
}

func (r *UsersRepository) Add(ctx context.Context, credentials *dtos.AuthCredentials) (int64, error) {
	query := `
	INSERT INTO users (login, password) 
	VALUES ($1, $2) 
	RETURNING id;`

	var userID int64
	err := r.db.Connection.QueryRowContext(ctx, query, credentials.Login, credentials.Password).Scan(&userID)
	if err == nil {
		return userID, err
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) || !pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
		return -1, err
	}

	return -1, services.ErrLoginAlreadyExists
}
