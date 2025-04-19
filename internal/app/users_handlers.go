package app

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"io"
	"main/internal/adapters"
	"main/internal/constants"
	"main/internal/dtos"
	"main/internal/interfaces"
	"main/internal/schemas"
	"main/internal/services"
	"net/http"
	"time"
)

func NewUsersHandler(s interfaces.UsersService) *UsersHandler {
	return &UsersHandler{
		service: s,
		logger:  adapters.GetLogger(),
	}
}

type UsersHandler struct {
	service interfaces.UsersService
	logger  *zap.SugaredLogger
}

func (h *UsersHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authCredentials := &schemas.AuthCredentials{}
	err = easyjson.Unmarshal(body, authCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	credentials := dtos.AuthCredentials(*authCredentials)

	userID, err := h.service.Register(ctx, &credentials)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrLoginAlreadyExists):
			w.WriteHeader(http.StatusConflict)
			return
		default:
			h.logger.Infow(
				"Register user",
				"error", err.Error(),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	cookie, err := setJWTCookie(userID)
	if err != nil {
		h.logger.Infow(
			"Set jwt cookie",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//200 — пользователь успешно зарегистрирован и аутентифицирован;
	//400 — неверный формат запроса;
	//409 — логин уже занят;
	//500 — внутренняя ошибка сервера.

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authCredentials := &schemas.AuthCredentials{}
	err = easyjson.Unmarshal(body, authCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	credentials := dtos.AuthCredentials(*authCredentials)

	userID, err := h.service.Login(ctx, &credentials)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrAuthCredentialsIsNotValid):
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			h.logger.Infow(
				"Login user",
				"error", err.Error(),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	cookie, err := setJWTCookie(userID)
	if err != nil {
		h.logger.Infow(
			"Set jwt cookie",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//200 — пользователь успешно аутентифицирован;
	//400 — неверный формат запроса;
	//401 — неверная пара логин/пароль;
	//500 — внутренняя ошибка сервера.

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func setJWTCookie(userID int64) (*http.Cookie, error) {
	tokenStr, err := generateJWT(userID)
	if err != nil {
		return nil, err
	}
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    tokenStr,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   constants.CookieMaxAge,
	}
	return &cookie, nil
}

func generateJWT(userID int64) (string, error) {
	expirationTime := time.Now().Add(constants.TokenExp)
	claims := &schemas.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(constants.JwtSecret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
