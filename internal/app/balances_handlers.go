package app

import (
	"errors"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
	"go.uber.org/zap"
	"io"
	"main/internal/adapters"
	"main/internal/constants"
	"main/internal/dtos"
	"main/internal/interfaces"
	"main/internal/schemas"
	"main/internal/services"
	"net/http"
)

func NewBalancesHandler(s interfaces.BalancesService) *BalancesHandler {
	return &BalancesHandler{
		service: s,
		logger:  adapters.GetLogger(),
	}
}

type BalancesHandler struct {
	service interfaces.BalancesService
	logger  *zap.SugaredLogger
}

func (h *BalancesHandler) Get(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.service.Get(ctx)
	if err != nil {
		h.logger.Infow(
			"Balance get",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := easyjson.Marshal(schemas.Balance(*result))
	if err != nil {
		h.logger.Infow(
			"Balance get Marshal",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", constants.JSONContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *BalancesHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
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

	withdraw := &schemas.Withdraw{}
	err = easyjson.Unmarshal(body, withdraw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	withdrawData := dtos.Withdraw(*withdraw)

	err = h.service.Withdraw(ctx, &withdrawData)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrPaymentRequired):
			w.WriteHeader(http.StatusPaymentRequired)
			return
		case errors.Is(err, services.ErrOrderIDNotValid):
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		default:
			h.logger.Infow(
				"Balance withdraw",
				"error", err.Error(),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BalancesHandler) Withdrawals(w http.ResponseWriter, r *http.Request) {
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

	results, err := h.service.Withdrawals(ctx)
	if err != nil {
		h.logger.Infow(
			"Balance withdrawals",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	items := make([]schemas.Withdrawal, len(results))
	for i := 0; i < len(results); i++ {
		items[i] = schemas.Withdrawal(*results[i])
	}

	var writer jwriter.Writer
	err = marshalWithdrawalSlice(items, &writer)
	if err != nil {
		h.logger.Infow(
			"Balance withdrawals Marshal",
			"error", err.Error(),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", constants.JSONContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(writer.Buffer.BuildBytes())
}

func marshalWithdrawalSlice(v []schemas.Withdrawal, wr *jwriter.Writer) error {
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
