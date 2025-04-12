package repositories

const (
	// Links
	addShortLink = `
		INSERT INTO events (short, origin, user_id) 
		VALUES ($1, $2, $3)`
	getShortLink = `
		SELECT origin, is_deleted 
		FROM events 
		WHERE short = $1;`
	getOrigin = `
		SELECT short 
		FROM events 
		WHERE origin = $1;`
	// Users
	addUser = `
		INSERT INTO users DEFAULT VALUES RETURNING id;`
	getLinksByUser = `
		SELECT short, origin FROM events WHERE user_id = $1;`
	deleteLinksByUser = `
		UPDATE events SET is_deleted = true WHERE short = $1 AND user_id = $2;`
)
