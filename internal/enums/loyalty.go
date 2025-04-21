package enums

type LoyaltyStatusEnum uint8

const (
	LoyaltyRegistered LoyaltyStatusEnum = iota + 1
	LoyaltyProcessing
	LoyaltyInvalid
	LoyaltyProcessed
)

var LoyaltyStatuses = map[string]LoyaltyStatusEnum{
	"REGISTERED": LoyaltyRegistered,
	"PROCESSING": LoyaltyProcessing,
	"INVALID":    LoyaltyInvalid,
	"PROCESSED":  LoyaltyProcessed,
}

func (o LoyaltyStatusEnum) String() string {
	for k, v := range LoyaltyStatuses {
		if o == v {
			return k
		}
	}
	return ""
}

func LoyaltyStatusFromString(s string) (LoyaltyStatusEnum, bool) {
	status, ok := LoyaltyStatuses[s]
	return status, ok
}
