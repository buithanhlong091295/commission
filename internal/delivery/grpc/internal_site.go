package grpc

import (
	"context"
	"xtek/exchange/commission/internal/domain/commission"
	pb "xtek/exchange/commission/pb/commission"
	pbDTO "xtek/exchange/commission/pb/commission/dto"
)

// NewInternalSiteDelivery ...
func NewInternalSiteDelivery(comDomain *commission.CommissionDomain) (pb.InternalSiteServiceServer, error) {
	return &internalSiteDelivery{comDomain}, nil
}

type internalSiteDelivery struct {
	comDomain *commission.CommissionDomain
}

func (s *internalSiteDelivery) GetTotalAmountFreezedByUserIDNCoin(ctx context.Context, req *pbDTO.GetTotalAmountFreezedByUserIDNCoinRequest) (*pbDTO.GetTotalAmountFreezedByUserIDNCoinResponse, error) {
	res, err := s.comDomain.GetTotalAmountFreezedByCoinNUserID(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
