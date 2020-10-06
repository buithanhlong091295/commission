package commission

import (
	"context"
	dErrors "xtek/exchange/commission/internal/errors"
	pbDTO "xtek/exchange/commission/pb/commission/dto"
	pbTypes "xtek/exchange/commission/pb/commission/types"
	pbUserDTO "xtek/exchange/commission/pb/user/dto"
)

// GetMembersBySponsorID ..
func (c *CommissionDomain) GetMembersBySponsorID(ctx context.Context, req *pbDTO.GetMembersByUserIDRequest) (*pbDTO.GetMembersByUserIDResponse, error) {
	users, err := c.userClient.GetMembersBySponsorID(ctx, &pbUserDTO.GetMembersBySponsorIDRequest{UserID: req.GetUserID()})
	if err != nil {
		return nil, dErrors.ErrCannotFindUsers
	}
	usInfo := []*pbTypes.Member{}
	for _, u := range users.GetListings() {
		usInfo = append(usInfo, &pbTypes.Member{
			Email:     u.Email,
			SponsorID: u.SponsorID,
			UserID:    u.UserID,
		})
	}
	res := &pbDTO.GetMembersByUserIDResponse{
		Users: usInfo,
	}
	return res, nil
}
