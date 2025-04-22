package schemas

import "github.com/golang-jwt/jwt/v4"

//easyjson -all internal/schemas/auth.go

// easyjson:skip
type Claims struct {
	UserID int64 `json:"userId"`
	jwt.RegisteredClaims
}
type AuthCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
