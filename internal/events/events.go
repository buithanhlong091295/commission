package events

import (
	"github.com/richard-xtek/go-grpc-micro-kit/kafka"
)

//
const (
	// ETTCommissionHistory ...
	ETTSaveCommissionHistory = kafka.EventType("commission.history.save")
)

// Define topic
const (
	CommissionTopic = "commission.response"
)

type kafkaTopic struct {
}

// Topic ...
var Topic kafkaTopic

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
