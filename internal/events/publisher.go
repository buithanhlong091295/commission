package events

// import (
// 	"context"
// 	"encoding/json"
// 	"xtek/exchange/commission/internal/models"
// 	pbContactTypes "xtek/exchange/commission/pb/contact/types"
// 	pbLendingTypes "xtek/exchange/commission/pb/lending/types"
// 	pbWalletTypes "xtek/exchange/commission/pb/wallet/types"

// 	"github.com/richard-xtek/go-grpc-micro-kit/kafka"

// 	ctx_logf "github.com/richard-xtek/go-grpc-micro-kit/grpc-logf/ctx-logf"
// 	"go.uber.org/zap"
// )

// // NewPublisher ...
// func NewPublisher(pub *kafka.Publisher) *Publisher {
// 	return &Publisher{pub}
// }

// // Publisher ...
// type Publisher struct {
// 	pub *kafka.Publisher
// }

// // PublishUserBalanceUpdated ...
// func (p *Publisher) PublishUserBalanceUpdated(ctx context.Context, userID string, balances []*models.BalanceUpdatedEvent) error {
// 	logger := ctx_logf.Extract(ctx)

// 	payloadBytes, err := json.Marshal(balances)
// 	if err != nil {
// 		logger.Bg().Error("Json marshal error", zap.Error(err))
// 		return err
// 	}

// 	msgWrapper := models.SocketEvent{
// 		Data:     payloadBytes,
// 		Platform: models.TraddingPlatformInternal,
// 		UserID:   userID,
// 	}
// 	marshaler := kafka.JSONMarshaler{}
// 	msg, err := marshaler.Marshal(msgWrapper)
// 	if err != nil {
// 		logger.Bg().Error("Json marshal error", zap.Error(err))
// 		return err
// 	}
// 	return p.pub.Publish(Topic.UserBalanceUpdatedEvent(), msg)
// }

// // PublishTransactionWithdrawConfirmed ...
// func (p *Publisher) PublishTransactionWithdrawConfirmed(ctx context.Context, transaction *models.Transaction) error {
// 	logger := ctx_logf.Extract(ctx)

// 	payload := transaction.MarshalEventProtobuf()

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(payload)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTTransactionWithdrawConfirmed.String())

// 	return p.pub.Publish(WalletEngineTopic, msg)
// }

// // PublishAdminApprovalWithdraw ...
// func (p *Publisher) PublishAdminApprovalWithdraw(ctx context.Context, transaction *models.Transaction) error {
// 	logger := ctx_logf.Extract(ctx)

// 	payload := transaction.MarshalEventProtobuf()

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(payload)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTAdminApprovalWithdraw.String())

// 	return p.pub.Publish(WithdrawApprovalEngineTopic, msg)
// }

// // PublishUserWithdrawSuccess ...
// func (p *Publisher) PublishUserWithdrawSuccess(ctx context.Context, transaction *models.Transaction) error {
// 	logger := ctx_logf.Extract(ctx)

// 	payload := &pbContactTypes.SendWithdrawSuccess{
// 		TxID:    transaction.TxID,
// 		Address: transaction.Address,
// 		Amount:  transaction.Amount.String(),
// 		Coin:    transaction.Coin,
// 		UserID:  transaction.UserID,
// 	}

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(payload)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTSendUserWithdrawSuccess.String())

// 	return p.pub.Publish(Topic.ContactEngine(), msg)
// }

// // PublishUserTransferPending ...
// func (p *Publisher) PublishUserTransferPending(ctx context.Context, transferHistory *models.TransferHistory) error {
// 	logger := ctx_logf.Extract(ctx)

// 	payload := &pbWalletTypes.Transfer{
// 		Amount: transferHistory.Amount.String(),
// 		Coin:   transferHistory.Coin,
// 		UserID: transferHistory.UserID,
// 		From:   models.BalanceTypeToProtobuf(transferHistory.From),
// 		To:     models.BalanceTypeToProtobuf(transferHistory.To),
// 		Id:     transferHistory.ID,
// 	}

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(payload)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTTransferPending.String())

// 	return p.pub.Publish(WalletEngineTopic, msg)
// }

// // PublishContractApplySuccessResponseEvent ...
// func (p *Publisher) PublishContractApplySuccessResponseEvent(ctx context.Context, contract *pbLendingTypes.Contract) error {
// 	logger := ctx_logf.Extract(ctx)

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(contract)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTContractApplyCompletedEvent.String())

// 	return p.pub.Publish(ContractResponseTopic, msg)
// }

// // PublishContractApplyFailureResponseEvent ...
// func (p *Publisher) PublishContractApplyFailureResponseEvent(ctx context.Context, contract *pbLendingTypes.Contract) error {
// 	logger := ctx_logf.Extract(ctx)

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(contract)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTContractApplyFailureEvent.String())

// 	return p.pub.Publish(ContractResponseTopic, msg)
// }

// // PublishOrderMatchingEngineResponse ...
// func (p *Publisher) PublishOrderMatchingEngineResponse(ctx context.Context, res *pbOrderTypes.OrderMatchingEngineResponse) error {
// 	logger := ctx_logf.Extract(ctx)

// 	marshaler := kafka.ProtobufMarshaler{}
// 	msg, err := marshaler.Marshal(res)
// 	if err != nil {
// 		logger.Bg().Error("Marshal", zap.Error(err))
// 	}

// 	msg.Metadata.Set(kafka.EventTypeHeaderKey, ETTOrderPlaceResponse.String())

// 	return p.pub.Publish(OrderEngineResponseTopic, msg)
// }
