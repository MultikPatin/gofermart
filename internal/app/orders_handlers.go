package app

import (
	"errors"
	"github.com/mailru/easyjson/jwriter"
	"go.uber.org/zap"
	"io"
	"main/internal/adapters"
	"main/internal/constants"
	"main/internal/interfaces"
	"main/internal/schemas"
	"main/internal/services"
	"net/http"
)

func NewOrdersHandler(s interfaces.OrdersService) *OrdersHandler {
	return &OrdersHandler{
		service: s,
		logger:  adapters.GetLogger(),
	}
}

type OrdersHandler struct {
	service interfaces.OrdersService
	logger  *zap.SugaredLogger
}

func (h *OrdersHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	UserAuth := ctx.Value(constants.UserAuth).(bool)
	if !UserAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	OrderID := string(body)

	err = h.service.Add(ctx, OrderID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrOrderAlreadyExists):
			w.WriteHeader(http.StatusOK)
			return
		case errors.Is(err, services.ErrOrderAlreadyLoadedByAnotherUser):
			w.WriteHeader(http.StatusConflict)
			return
		case errors.Is(err, services.ErrOrderIDNotValid):
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		default:
			h.logger.Infow(
				"Order add",
				"error", err.Error(),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *OrdersHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	UserAuth := ctx.Value(constants.UserAuth).(bool)
	if !UserAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	results, err := h.service.GetAll(ctx)
	if err != nil {
		h.logger.Infow(
			"Order get all",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	items := make([]schemas.Order, 0, len(results))
	for i := 0; i < len(results); i++ {
		item := schemas.Order{
			Number:   results[i].Number,
			Status:   results[i].Status.String(),
			Accrual:  results[i].Accrual,
			Uploaded: results[i].Uploaded,
		}
		items = append(items, item)
	}

	var writer jwriter.Writer
	err = marshalOrderSlice(items, &writer)
	if err != nil {
		h.logger.Infow(
			"Balance get all Marshal",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", constants.JSONContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(writer.Buffer.BuildBytes())
}

func marshalOrderSlice(v []schemas.Order, wr *jwriter.Writer) error {
	wr.RawByte('[')
	for i := 0; i < len(v); i++ {
		if i > 0 {
			wr.RawByte(',')
		}
		v[i].MarshalEasyJSON(wr)
	}
	wr.RawByte(']')
	return nil
}
