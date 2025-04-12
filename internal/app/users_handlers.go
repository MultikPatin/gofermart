package app

import (
	"github.com/mailru/easyjson"
	"io"
	"main/internal/dtos"
	"main/internal/interfaces"
	"main/internal/schemas"
	"net/http"
)

func NewUsersHandler(s interfaces.UsersService) *UsersHandler {
	return &UsersHandler{
		service: s,
	}
}

type UsersHandler struct {
	service interfaces.UsersService
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
	err = h.service.Register(ctx, dtos.AuthCredentials(*authCredentials))
	if err != nil {
		// add log
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//200 — пользователь успешно зарегистрирован и аутентифицирован;
	//400 — неверный формат запроса;
	//409 — логин уже занят;
	//500 — внутренняя ошибка сервера.

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

	err = h.service.Login(ctx, dtos.AuthCredentials(*authCredentials))
	if err != nil {
		// add log
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//200 — пользователь успешно аутентифицирован;
	//400 — неверный формат запроса;
	//401 — неверная пара логин/пароль;
	//500 — внутренняя ошибка сервера.

	w.WriteHeader(http.StatusOK)
}
