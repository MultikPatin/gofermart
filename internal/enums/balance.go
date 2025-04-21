package enums

type BalanceActionsEnum uint8

const (
	BalanceWithdrawal BalanceActionsEnum = iota + 1
	BalanceDeposit
)

var balanceActions = map[string]BalanceActionsEnum{
	"DEPOSIT":    BalanceWithdrawal,
	"WITHDRAWAL": BalanceDeposit,
}

func (o BalanceActionsEnum) String() string {
	for k, v := range balanceActions {
		if o == v {
			return k
		}
	}
	return ""
}

func BalanceActionsFromString(s string) (OrderStatusEnum, bool) {
	status, ok := orderStatuses[s]
	return status, ok
}
