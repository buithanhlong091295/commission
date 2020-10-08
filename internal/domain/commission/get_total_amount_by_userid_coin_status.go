package commission

import (
	"context"
	"xtek/exchange/commission/internal/models"
	pbDTO "xtek/exchange/commission/pb/commission/dto"

	ctx_logf "github.com/richard-xtek/go-grpc-micro-kit/grpc-logf/ctx-logf"
	"go.uber.org/zap"
)

// GetTotalAmountFreezedByCoinNUserID ..
func (c *CommissionDomain) GetTotalAmountFreezedByCoinNUserID(ctx context.Context, req *pbDTO.GetTotalAmountFreezedByUserIDNCoinRequest) (*pbDTO.GetTotalAmountFreezedByUserIDNCoinResponse, error) {
	logger := ctx_logf.Extract(ctx)
	logFields := []zap.Field{
		zap.String("func_domain", "GetTotalAmountFreezedByCoinNUserID"),
	}
	total, err := c.commissionRepo.GetTotalAmountByUserIDNCoinNStatus(req.GetCoin(), req.GetUserID(), models.CommissionFreezedStatus)
	if err != nil {
		logger.Bg().Error("GetTotalAmountByUserIDNCoinNStatus", append(logFields, zap.Error(err))...)
		return nil, err
	}
	res := &pbDTO.GetTotalAmountFreezedByUserIDNCoinResponse{
		Amount: total.String(),
	}
	return res, nil
}
