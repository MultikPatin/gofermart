package postgres

const (
	createTables = `
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
		    login VARCHAR(255) UNIQUE NOT NULL,
		    password VARCHAR(255) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS orders (
		    id SERIAL PRIMARY KEY,
		    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		    order_id VARCHAR(255) UNIQUE NOT NULL,
		    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS balances (
		    id SERIAL PRIMARY KEY,
		    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		    processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		    order_id VARCHAR(255) NOT NULL,
		    action VARCHAR(32) NOT NULL,
		    sum INTEGER NOT NULL
		);
		`
)
