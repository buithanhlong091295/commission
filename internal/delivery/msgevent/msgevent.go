package msgevent

import "xtek/exchange/commission/internal/domain/commission"

// New ...
func New(
	commssionDomain *commission.CommissionDomain,
) *MsgEvent {
	return &MsgEvent{
		commssionDomain,
	}
}

// MsgEvent ...
type MsgEvent struct {
	commssionDomain *commission.CommissionDomain
}
