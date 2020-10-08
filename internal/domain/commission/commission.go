package commission

import (
	"xtek/exchange/commission/internal/stores"
	pbUser "xtek/exchange/commission/pb/user"
)

// NewCommissionDomain create new user controller
func NewCommissionDomain(
	userClient pbUser.InternalSiteServiceClient,
	commissionRepo stores.CommissionRepo,
) *CommissionDomain {
	return &CommissionDomain{
		userClient,
		commissionRepo,
	}
}

// CommissionDomain struct
type CommissionDomain struct {
	userClient     pbUser.InternalSiteServiceClient
	commissionRepo stores.CommissionRepo
}
