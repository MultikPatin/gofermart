package app

import (
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
	"io"
	"main/internal/constants"
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

	//http.StatusConflict
	//http.StatusInternalServerError
	//http.StatusBadRequest
	//http.StatusOK

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

	//http.StatusConflict
	//http.StatusInternalServerError
	//http.StatusBadRequest
	//http.StatusOK

	w.WriteHeader(http.StatusOK)
}

func (h *UsersHandler) Withdrawals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	results, err := h.service.Withdrawals(ctx)
	if err != nil {
		// add log
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	items := make([]schemas.Withdrawal, 0, len(results))
	for i := 0; i < len(results); i++ {
		items = append(items, schemas.Withdrawal(*results[i]))
	}

	var writer jwriter.Writer
	err = MarshalSliceEasyJSON(items, &writer)
	if err != nil {
		// add log
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//http.StatusOK
	//http.StatusNoContent
	//http.StatusUnauthorized
	//http.StatusInternalServerError

	w.Header().Set("content-type", constants.JSONContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(writer.Buffer.BuildBytes())
}

func MarshalSliceEasyJSON(v []schemas.Withdrawal, wr *jwriter.Writer) error {
	wr.RawByte('[')
	for i, w := range v {
		if i > 0 {
			wr.RawByte(',')
		}
		w.MarshalEasyJSON(wr)
	}
	wr.RawByte(']')
	return nil
}
