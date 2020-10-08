package msgevent

import (
	"context"
	"errors"

	pbLendingTypes "xtek/exchange/commission/pb/lending/types"

	ctx_logf "github.com/richard-xtek/go-grpc-micro-kit/grpc-logf/ctx-logf"
	"go.uber.org/zap"
)

// CommissionEvent ...
func (m *MsgEvent) CommissionEvent(ctx context.Context, payload interface{}) error {
	logger := ctx_logf.Extract(ctx)
	data, ok := payload.(*pbLendingTypes.CommissionEvent)
	if !ok {
		logger.Bg().Error("Unknow data type for CommissionEvent", zap.String("function", "CommissionEvent"))
		return errors.New("")
	}

	return m.commssionDomain.CreateCommissionWithEvent(ctx, data)
}
