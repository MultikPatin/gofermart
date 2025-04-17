package middleware

import (
	"context"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"main/internal/constants"
	"main/internal/schemas"
	"net/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil || cookie == nil {
			userID := -1
			ctx := context.WithValue(r.Context(), constants.UserIDKey, userID)
			ctx = context.WithValue(ctx, constants.UserAuth, false)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			tokenStr := cookie.Value
			claims, err := verifyJWT(tokenStr)
			if err != nil {
				userID := -1
				ctx := context.WithValue(r.Context(), constants.UserIDKey, userID)
				ctx = context.WithValue(ctx, constants.UserAuth, false)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
			ctx := context.WithValue(r.Context(), constants.UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, constants.UserAuth, true)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func verifyJWT(tokenStr string) (*schemas.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &schemas.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*schemas.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
