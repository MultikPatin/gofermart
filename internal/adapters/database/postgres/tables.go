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
		    accrual FLOAT NOT NULL,
			status VARCHAR(255) NOT NULL,
		    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS balances (
		    id SERIAL PRIMARY KEY,
		    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		    order_id VARCHAR(255),
		    processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		    action VARCHAR(32) NOT NULL,
		    amount FLOAT NOT NULL CHECK (amount >= 0)
		);
		CREATE INDEX IF NOT EXISTS order_id_index ON balances(order_id);
		DO
		
		$$
		BEGIN
			CREATE TYPE balance_actions_enum AS ENUM ('DEPOSIT', 'WITHDRAWAL');
		EXCEPTION WHEN duplicate_object THEN
			RAISE NOTICE 'Тип данных balance_actions_enum уже существует.';
		END
		
		$$;
		ALTER TABLE balances ALTER COLUMN action TYPE balance_actions_enum USING action::balance_actions_enum;
		DO
		
		$$
		BEGIN
			CREATE TYPE order_status_enum AS ENUM ('NEW', 'PROCESSING', 'INVALID', 'PROCESSED');
		EXCEPTION WHEN duplicate_object THEN
			RAISE NOTICE 'Тип данных order_status_enum уже существует.';
		END
		
		$$;
		ALTER TABLE orders ALTER COLUMN status TYPE order_status_enum USING status::order_status_enum;`
)
