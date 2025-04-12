package app

import (
	"errors"
	"main/internal/constants"
	"main/internal/interfaces"
	"main/internal/services"
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

	originLink, err := h.service.Register(ctx, r.PathValue("id"))
	if err != nil {
		if errors.Is(err, services.ErrDeletedLink) {
			http.Error(w, "Origin is deleted", http.StatusGone)
		} else {
			http.Error(w, "Origin not found", http.StatusNotFound)
		}
		return
	}

	w.Header().Set("content-type", constants.TextContentType)
	w.Header().Set("Location", originLink)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	originLink, err := h.service.Login(ctx, r.PathValue("id"))
	if err != nil {
		if errors.Is(err, services.ErrDeletedLink) {
			http.Error(w, "Origin is deleted", http.StatusGone)
		} else {
			http.Error(w, "Origin not found", http.StatusNotFound)
		}
		return
	}

	w.Header().Set("content-type", constants.TextContentType)
	w.Header().Set("Location", originLink)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *UsersHandler) Withdrawals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	originLink, err := h.service.Withdrawals(ctx, r.PathValue("id"))
	if err != nil {
		if errors.Is(err, services.ErrDeletedLink) {
			http.Error(w, "Origin is deleted", http.StatusGone)
		} else {
			http.Error(w, "Origin not found", http.StatusNotFound)
		}
		return
	}

	w.Header().Set("content-type", constants.TextContentType)
	w.Header().Set("Location", originLink)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
