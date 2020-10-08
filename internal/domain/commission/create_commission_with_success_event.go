package commission

import (
	"context"
	"xtek/exchange/commission/internal/models"
	pbLendingTypes "xtek/exchange/commission/pb/lending/types"

	ctx_logf "github.com/richard-xtek/go-grpc-micro-kit/grpc-logf/ctx-logf"
	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"
	"go.uber.org/zap"
)

// CreateCommissionWithEvent ..
func (c *CommissionDomain) CreateCommissionWithEvent(ctx context.Context, payload *pbLendingTypes.CommissionEvent) error {
	logger := ctx_logf.Extract(ctx)
	logFields := []zap.Field{
		zap.String("func_domain", "CreateCommissionWithEvent"),
	}
	com := &models.CommissionHistory{
		Status:     models.CommissionStatus(payload.GetStatus()),
		SenderID:   payload.GetSenderID(),
		ReceiverID: payload.GetReceiverID(),
		Coin:       payload.GetCoin(),
		Type:       models.CommissionType(payload.GetType()),
		Amount:     math.ToBigInt(payload.GetAmount()),
	}
	err := c.commissionRepo.Create(com)
	if err != nil {
		logger.Bg().Error("commissionRepo.Create", append(logFields, zap.Error(err))...)
		return err
	}
	return nil
}
