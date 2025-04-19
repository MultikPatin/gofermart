package enums

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

func OrdesStatusFromString(s string) (OrderStatusEnum, bool) {
	status, ok := orderStatuses[s]
	return status, ok
}
