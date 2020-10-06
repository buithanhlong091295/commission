package grpc

import (
	"context"
	pb "xtek/exchange/commission/pb/commission"
	pbDTO "xtek/exchange/commission/pb/commission/dto"
	pbTypes "xtek/exchange/commission/pb/commission/types"
	pbPaginate "xtek/exchange/commission/pb/paginate"
)

// NewUserSiteDelivery ...
func NewUserSiteDelivery() (pb.UserSiteServiceServer, error) {
	return &userSiteDelivery{}, nil
}

type userSiteDelivery struct {
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
			CreatedAt:   1601958158923,
		},
		&pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601958158923,
		}, &pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601958158923,
		}, &pbTypes.Commission{
			Id:          "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			SenderID:    "bbbbbbbbbbbbbbbbbbbbbbbbbbb",
			SenderEmail: "abc@gmail.com",
			ReceiverID:  "cccccccccccccccccccccccccccc",
			Amount:      "123",
			Status:      1,
			CreatedAt:   1601958158923,
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

	var listings = []*pbTypes.Member{
		&pbTypes.Member{
			UserID:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Email:      "abc@gmail.com",
			Earned:     "12345",
			NumMembers: 123,
			SponsorID:  "ddddddddddddddddddddddddddddd",
		},
		&pbTypes.Member{
			UserID:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Email:      "abc@gmail.com",
			Earned:     "12345",
			NumMembers: 123,
			SponsorID:  "ddddddddddddddddddddddddddddd",
		},
		&pbTypes.Member{
			UserID:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Email:      "abc@gmail.com",
			Earned:     "12345",
			NumMembers: 123,
			SponsorID:  "ddddddddddddddddddddddddddddd",
		},
		&pbTypes.Member{
			UserID:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Email:      "abc@gmail.com",
			Earned:     "12345",
			NumMembers: 123,
			SponsorID:  "ddddddddddddddddddddddddddddd",
		},
	}

	res := &pbDTO.GetMembersByUserIDResponse{
		Users: listings,
	}

	return res, nil
}
