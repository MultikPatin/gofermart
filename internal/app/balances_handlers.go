package app

import (
	"errors"
	"main/internal/constants"
	"main/internal/interfaces"
	"main/internal/services"
	"net/http"
)

func NewBalancesHandler(s interfaces.BalancesService) *BalancesHandler {
	return &BalancesHandler{
		service: s,
	}
}

type BalancesHandler struct {
	service interfaces.BalancesService
}

func (h *BalancesHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	originLink, err := h.service.Get(ctx, r.PathValue("id"))
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

func (h *BalancesHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	originLink, err := h.service.Withdraw(ctx, r.PathValue("id"))
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
