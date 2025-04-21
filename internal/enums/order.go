package enums

import "fmt"

type OrderStatusEnum uint8

const (
	OrderCreated OrderStatusEnum = iota + 1
	OrderProcessing
	OrderInvalid
	OrderProcessed
)

var orderStatuses = map[string]OrderStatusEnum{
	"NEW":        OrderCreated,
	"PROCESSING": OrderProcessing,
	"INVALID":    OrderInvalid,
	"PROCESSED":  OrderProcessed,
}

func (o OrderStatusEnum) String() string {
	for k, v := range orderStatuses {
		if o == v {
			return k
		}
	}
	return ""
}

func OrdersStatusFromString(s string) (OrderStatusEnum, bool) {
	status, ok := orderStatuses[s]
	return status, ok
}

func MutateLoyaltyToOrderStatus(status LoyaltyStatusEnum) (OrderStatusEnum, error) {
	switch status {
	case LoyaltyRegistered:
		return OrderCreated, nil
	case LoyaltyProcessing:
		return OrderProcessing, nil
	case LoyaltyInvalid:
		return OrderInvalid, nil
	case LoyaltyProcessed:
		return OrderProcessed, nil
	default:
		return 0, fmt.Errorf("invalid loyalty status")
	}
}
