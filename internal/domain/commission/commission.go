package commission

import (
	"xtek/exchange/commission/internal/stores"
	pbUser "xtek/exchange/commission/pb/user"
)

// NewCommissionDomain create new user controller
func NewCommissionDomain(
	userClient pbUser.InternalSiteServiceClient,
	commissionRepo stores.CommissionRepo,
	usersMemberRepo stores.UsersMemberRepo,
) *CommissionDomain {
	return &CommissionDomain{
		userClient,
		commissionRepo,
		usersMemberRepo,
	}
}

// CommissionDomain struct
type CommissionDomain struct {
	userClient      pbUser.InternalSiteServiceClient
	commissionRepo  stores.CommissionRepo
	usersMemberRepo stores.UsersMemberRepo
}
