package grpc

import (
	"context"
	"xtek/exchange/commission/internal/domain/commission"
	pb "xtek/exchange/commission/pb/commission"
	pbDTO "xtek/exchange/commission/pb/commission/dto"
	pbTypes "xtek/exchange/commission/pb/commission/types"
	pbPaginate "xtek/exchange/commission/pb/paginate"
)

// NewUserSiteDelivery ...
func NewUserSiteDelivery(comDomain *commission.CommissionDomain) (pb.UserSiteServiceServer, error) {
	return &userSiteDelivery{comDomain: comDomain}, nil
}

type userSiteDelivery struct {
	comDomain *commission.CommissionDomain
}

func (s *userSiteDelivery) GetUserCommissions(ctx context.Context, req *pbDTO.GetUserCommissionsRequest) (*pbDTO.GetUserCommissionsResponse, error) {

	var listings = []*pbTypes.Commission{
		&pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601972004,
		},
		&pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601972004,
		}, &pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601972004,
		}, &pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601972004,
		},
	}
	paginate := &pbPaginate.PaginateResponse{
		Limit:       4,
		CurrentPage: 1,
		TotalPages:  1,
		Total:       4,
	}

	res := &pbDTO.GetUserCommissionsResponse{
		Listings:   listings,
		Pagination: paginate,
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
