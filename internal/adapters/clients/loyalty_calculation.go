package clients

import (
	"context"
	"fmt"
	"github.com/mailru/easyjson"
	"io"
	"main/internal/dtos"
	"main/internal/schemas"
	"net/http"
	"strings"
	"time"
)

type LoyaltyCalculation struct {
	accrualSystemAddr string
	client            *http.Client
}

func NewLoyaltyCalculation(Addr string) *LoyaltyCalculation {
	timeout := 3 * time.Second

	return &LoyaltyCalculation{
		accrualSystemAddr: Addr,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (l *LoyaltyCalculation) GetByOrderID(ctx context.Context, orderID string) (*dtos.LoyaltyCalculation, error) {
	endpoint := fmt.Sprintf("%s/api/orders/%s", l.accrualSystemAddr, orderID)

	request, err := http.NewRequest(http.MethodGet, endpoint, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	response, err := l.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	loyalty := &schemas.LoyaltyCalculation{}
	err = easyjson.Unmarshal(body, loyalty)
	if err != nil {
		return nil, err
	}

	result := dtos.LoyaltyCalculation{
		OrderBase: dtos.OrderBase{
			Number: loyalty.Number,
		},
		OrderStatus: dtos.OrderStatus{
			Status:  loyalty.Status,
			Accrual: loyalty.Accrual,
		},
	}

	return &result, nil
}
