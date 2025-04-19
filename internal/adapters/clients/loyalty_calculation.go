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
)

type LoyaltyCalculation struct {
	accrualSystemAddr string
}

func NewLoyaltyCalculation(Addr string) *LoyaltyCalculation {
	return &LoyaltyCalculation{
		accrualSystemAddr: Addr,
	}
}

func (l *LoyaltyCalculation) GetByOrderID(ctx context.Context, orderID string) (*dtos.LoyaltyCalculation, error) {
	endpoint := fmt.Sprintf("%s/api/orders/%s", l.accrualSystemAddr, orderID)
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, endpoint, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
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
