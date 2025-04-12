package postgres

const (
	createTables = `
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY
		    login VARCHAR(255) UNIQUE NOT NULL
		    password VARCHAR(255) NOT NULL
		);
		CREATE UNIQUE INDEX IF NOT EXISTS user_login_index ON users(login);
		`
)
