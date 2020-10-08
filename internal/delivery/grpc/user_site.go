package grpc

import (
	"context"
	"xtek/exchange/commission/internal/domain/commission"
	dErrors "xtek/exchange/commission/internal/errors"
	pb "xtek/exchange/commission/pb/commission"
	pbDTO "xtek/exchange/commission/pb/commission/dto"

	"github.com/richard-xtek/go-grpc-micro-kit/auth/requestinfo"
)

// NewUserSiteDelivery ...
func NewUserSiteDelivery(comDomain *commission.CommissionDomain) (pb.UserSiteServiceServer, error) {
	return &userSiteDelivery{comDomain: comDomain}, nil
}

type userSiteDelivery struct {
	comDomain *commission.CommissionDomain
}

func (s *userSiteDelivery) GetUserCommissions(ctx context.Context, req *pbDTO.GetUserCommissionsRequest) (*pbDTO.GetUserCommissionsResponse, error) {
	claim, isOk := requestinfo.ExtractRequestInfo(ctx)
	if !isOk {
		return nil, dErrors.ErrAuthenticationFail
	}

	if claim.UserID == "" {
		return nil, dErrors.ErrAuthenticationFail
	}
	res, err := s.comDomain.GetCommissionsByFilter(ctx, req, claim.UserID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userSiteDelivery) GetMembersByUserID(ctx context.Context, req *pbDTO.GetMembersByUserIDRequest) (*pbDTO.GetMembersByUserIDResponse, error) {
	res, err := s.comDomain.GetMembersBySponsorID(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
