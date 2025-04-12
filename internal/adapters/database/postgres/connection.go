package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"main/internal/adapters"
	"net/url"
	"time"
)

type Database struct {
	Connection *sql.DB
}

func (p *Database) Close() error {
	err := p.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *Database) Ping() error {
	err := p.Connection.Ping()
	return err
}

func NewDatabase(DNS *url.URL) (*Database, error) {
	logger := adapters.GetLogger()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	host := DNS.Hostname()
	port := DNS.Port()
	user := DNS.User.Username()
	password, _ := DNS.User.Password()
	dbname := DNS.Path[1:]

	ps := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("pgx", ps)
	if err != nil {
		logger.Infow(
			"Create Postgres Database",
			"error", err.Error(),
		)
	}

	err = migrate(ctx, conn)
	if err != nil {
		logger.Infow(
			"Create tables",
			"error", err.Error(),
		)
	}
	postgresDB := Database{
		Connection: conn,
	}
	return &postgresDB, err
}

func migrate(ctx context.Context, conn *sql.DB) error {
	_, err := conn.ExecContext(ctx, createTables)
	if err != nil {
		return err
	}
	return nil
}
