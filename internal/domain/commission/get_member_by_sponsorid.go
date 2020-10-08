package commission

import (
	"context"
	dErrors "xtek/exchange/commission/internal/errors"
	pbDTO "xtek/exchange/commission/pb/commission/dto"
	pbTypes "xtek/exchange/commission/pb/commission/types"
	pbUserDTO "xtek/exchange/commission/pb/user/dto"

	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"
)

// GetMembersBySponsorID ..
func (c *CommissionDomain) GetMembersBySponsorID(ctx context.Context, req *pbDTO.GetMembersByUserIDRequest) (*pbDTO.GetMembersByUserIDResponse, error) {
	users, err := c.userClient.GetMembersBySponsorID(ctx, &pbUserDTO.GetMembersBySponsorIDRequest{UserID: req.GetUserID()})
	if err != nil {
		return nil, err
	}
	usInfo := []*pbTypes.Member{}
	for _, u := range users.GetListings() {
		totalEarned, err := c.commissionRepo.GetTotalAmountEarnByUserID(u.GetUserID())
		if err != nil {
			return nil, err
		}
		usersMember, err := c.usersMemberRepo.GetByUserID(u.GetUserID())
		if err != nil {
			return nil, err
		}
		if usersMember == nil {
			return nil, dErrors.ErrUsersMemberNotFound
		}
		usInfo = append(usInfo, &pbTypes.Member{
			Email:      u.Email,
			SponsorID:  u.SponsorID,
			UserID:     u.UserID,
			EarnedUSDT: math.ToDecimalString(totalEarned),
			NumMembers: usersMember.TotalMembers,
		})
	}
	res := &pbDTO.GetMembersByUserIDResponse{
		Listings: usInfo,
	}
	return res, nil
}
