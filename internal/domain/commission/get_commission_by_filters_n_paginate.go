package commission

import (
	"context"
	"strconv"
	"xtek/exchange/commission/internal/models"
	pbDTO "xtek/exchange/commission/pb/commission/dto"
	pbTypes "xtek/exchange/commission/pb/commission/types"
	pbPaginate "xtek/exchange/commission/pb/paginate"
	pbUserDTO "xtek/exchange/commission/pb/user/dto"

	ctx_logf "github.com/richard-xtek/go-grpc-micro-kit/grpc-logf/ctx-logf"
	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"
	"go.uber.org/zap"
)

//GetCommissionsByFilter ...
func (c *CommissionDomain) GetCommissionsByFilter(ctx context.Context, req *pbDTO.GetUserCommissionsRequest, userID string) (*pbDTO.GetUserCommissionsResponse, error) {
	logger := ctx_logf.Extract(ctx)
	logFields := []zap.Field{
		zap.String("Status", strconv.Itoa(int(req.GetFilters().GetStatus()))),
		zap.String("Type", strconv.Itoa(int(req.GetFilters().GetComType()))),
		zap.String("Limit", strconv.Itoa(int(req.GetPagination().GetLimit()))),
		zap.String("Page", strconv.Itoa(int(req.GetPagination().GetPage()))),
		zap.String("func_domain", "GetCommissionsByFilter"),
	}
	query := c.makeQueryForGetCommission(req, userID)
	limit := int(req.GetPagination().GetLimit())
	page := int(req.GetPagination().GetPage())
	coms, err := c.commissionRepo.GetAllWithPaginateNFilter(query, limit, page)
	if err != nil {
		logger.Bg().Error("GetAllWithPaginateNFilter", append(logFields, zap.Error(err))...)
		return nil, err
	}
	res, err := c.convertComsToGetUserCommissionResponse(ctx, coms)
	if err != nil {
		logger.Bg().Error("convertComsToGetUserCommissionResponse", append(logFields, zap.Error(err))...)
		return nil, err
	}
	count, err := c.commissionRepo.CountByQuery(query)
	if err != nil {
		logger.Bg().Error("CountByQuery", append(logFields, zap.Error(err))...)
		return nil, err
	}
	totalPage := int(count / limit)
	if count%limit != 0 {
		totalPage++
	}

	// pagination
	pagination := &pbPaginate.PaginateResponse{
		Limit:       int32(limit),
		CurrentPage: int32(page),
		Total:       int32(count),
		TotalPages:  int32(totalPage),
	}
	res.Pagination = pagination
	return res, nil
}

func (c *CommissionDomain) convertComsToGetUserCommissionResponse(ctx context.Context, coms []*models.CommissionHistory) (*pbDTO.GetUserCommissionsResponse, error) {
	listings := []*pbTypes.Commission{}
	for _, com := range coms {
		user, err := c.userClient.GetUserEmailByID(ctx, &pbUserDTO.GetUserEmailByIDRequest{UserID: com.SenderID})
		if err != nil {
			return nil, err
		}
		detail := &pbTypes.Commission{
			Id:          com.ID,
			Amount:      math.ToDecimalString(com.Amount),
			SenderID:    com.SenderID,
			SenderEmail: user.GetEmail(),
			ReceiverID:  com.ReceiverID,
			Status:      int32(com.Status),
			CreatedAt:   com.CreatedAt.Unix(),
		}
		listings = append(listings, detail)
	}

	return &pbDTO.GetUserCommissionsResponse{Listings: listings}, nil
}

// makeQueryForGetCommission ...
func (c *CommissionDomain) makeQueryForGetCommission(req *pbDTO.GetUserCommissionsRequest, userID string) map[string]interface{} {
	comType := req.GetFilters().GetComType()
	status := req.GetFilters().GetStatus()
	query := make(map[string]interface{})
	query["receiverID"] = userID
	if comType == 0 || comType == 1 {
		query["type"] = comType
	}
	if status == 0 || status == 1 || status == 2 {
		query["status"] = status
	}

	return query
}
