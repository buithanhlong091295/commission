package commission

import (
	pbUser "xtek/exchange/commission/pb/user"
)

// NewCommissionDomain create new user controller
func NewCommissionDomain(
	userClient pbUser.InternalSiteServiceClient,
) *CommissionDomain {
	return &CommissionDomain{
		userClient,
	}
}

// CommissionDomain struct
type CommissionDomain struct {
	userClient pbUser.InternalSiteServiceClient
}
