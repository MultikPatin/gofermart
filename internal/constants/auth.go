package constants

import "time"

type userIDKey string
type userAuth bool

const (
	TokenExp               = time.Hour * 3
	CookieMaxAge           = 3600
	JwtSecret              = "your_secret_key"
	UserIDKey    userIDKey = "UserID"
	UserAuth     userAuth  = false
)
