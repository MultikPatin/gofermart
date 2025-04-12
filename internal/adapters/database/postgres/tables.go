package postgres

const (
	createTables = `
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY
		);
		CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		origin VARCHAR(255) NOT NULL UNIQUE,
		short VARCHAR(255) NOT NULL,
		is_deleted BOOLEAN DEFAULT FALSE);
		CREATE INDEX IF NOT EXISTS origin_index ON events(origin);`
)
