package events

// import (
// 	"fmt"

// 	"github.com/richard-xtek/go-grpc-micro-kit/kafka"
// )

// //
// const (
// 	// ETTTransactionWithdrawConfirmed ...
// 	ETTTransactionWithdrawConfirmed = kafka.EventType("transaction.withdraw.confirmed")
// 	// ETTAdminApprovalWithdraw use when bot publish, admin click approve
// 	ETTAdminApprovalWithdraw = kafka.EventType("transaction.withdraw.admin.approval")
// 	// ETTOrderPlaceRequest ...
// 	ETTOrderPlaceRequest = kafka.EventType("order.place.request")
// 	// ETTOrderPlaceResponse ...
// 	ETTOrderPlaceResponse = kafka.EventType("order.place.response")

// 	// ETTOrderMatchingEngine ...
// 	ETTOrderMatchingEngine = kafka.EventType("order.engine")

// 	// ETTTransactionBitcoinUnConfirmed ...
// 	ETTTransactionBitcoinUnConfirmed = kafka.EventType("transaction.bitcoin.unconfirmed")
// 	// ETTTransactionBitcoinUnConfirmed ...
// 	ETTTransactionBitcoinConfirmed = kafka.EventType("transaction.bitcoin.confirmed")
// 	// ETTSendUserWithdrawSuccess ...
// 	ETTSendUserWithdrawSuccess = kafka.EventType("contact.email.withdraw.success")

// 	// ETTTransactionDepositPending ...
// 	ETTTransactionDepositPending = kafka.EventType("transaction.deposit.pending")
// 	// ETTTransactionDepositCompleted ...
// 	ETTTransactionDepositCompleted = kafka.EventType("transaction.deposit.completed")
// 	// ETTTransferPending ...
// 	ETTTransferPending = kafka.EventType("transfer.pending")
// 	// ETTContractApplyCheckoutRequestEvent ...
// 	ETTContractApplyCheckoutRequestEvent = kafka.EventType("contract.apply.checkout.request")

// 	// ETTContractApplyCompletedEvent ...
// 	ETTContractApplyCompletedEvent = kafka.EventType("contract.apply.completed")
// 	// ETTContractApplyFailureEvent ...
// 	ETTContractApplyFailureEvent = kafka.EventType("contract.apply.failure")
// )

// // Define topic
// const (
// 	OrderEngineTopic            = "order.engine"
// 	OrderEngineResponseTopic    = "order.engine.response"
// 	WalletEngineTopic           = "wallet.engine"
// 	WithdrawApprovalEngineTopic = "withdraw.approval.engine"
// 	ContractResponseTopic       = "contract.response"
// )

// type kafkaTopic struct {
// }

// // Topic ...
// var Topic kafkaTopic

// func (t kafkaTopic) Kline(symbol string) string {
// 	return fmt.Sprintf("market.%s.kline.evt", symbol)
// }

// func (t kafkaTopic) MarketDepth(symbol string) string {
// 	return fmt.Sprintf("market.%s.depth.evt", symbol)
// }

// func (t kafkaTopic) MarketDetail(symbol string) string {
// 	return fmt.Sprintf("market.%s.detai.evt", symbol)
// }

// func (t kafkaTopic) MarketTradeDetail(symbol string) string {
// 	return fmt.Sprintf("market.%s.trade.detai.evt", symbol)
// }

// // OrderCreation ...
// func (t kafkaTopic) OrderCreation() string {
// 	return "order.create.evt"
// }

// // OrderMatching ...
// func (t kafkaTopic) OrderMatching() string {
// 	return "order.matching.evt"
// }

// // OrderCancellation ...
// func (t kafkaTopic) OrderCancellation() string {
// 	return "order.cancellation.evt"
// }
// func (t kafkaTopic) OrderEngineResponse() string {
// 	return "order.engine.response"
// }

// func (t kafkaTopic) OrderEngine() string {
// 	return "order.engine"
// }
// func (t kafkaTopic) UserEvent() string {
// 	return "user.evt"
// }

// func (t kafkaTopic) UserBalanceUpdatedEvent() string {
// 	return "user.balance.updated.evt"
// }

// func (t kafkaTopic) TransactionPending() string {
// 	return fmt.Sprintf("transaction.pending")
// }

// func (t *kafkaTopic) ContactEngine() string {
// 	return "contact.engine"
// }
