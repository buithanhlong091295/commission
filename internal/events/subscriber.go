package events

// import (
// 	"context"

// 	"github.com/richard-xtek/go-grpc-micro-kit/kafka"
// 	"github.com/richard-xtek/go-grpc-micro-kit/subscriber"

// 	"github.com/richard-xtek/go-grpc-micro-kit/log"
// 	"go.uber.org/zap"
// )

// // NewSubscriber ...
// func NewSubscriber(logger log.Factory, core *kafka.Subscriber, topic string) *Subscriber {
// 	return &Subscriber{
// 		logger: logger,
// 		core:   core,
// 		topic:  topic,
// 	}
// }

// // Subscriber ...
// type Subscriber struct {
// 	logger log.Factory
// 	topic  string

// 	core *kafka.Subscriber

// 	msgChan <-chan *kafka.Message

// 	ctx context.Context

// 	cancelCtxFn context.CancelFunc
// }

// // Start ...
// func (s *Subscriber) Start() error {
// 	s.ctx, s.cancelCtxFn = context.WithCancel(context.Background())

// 	msgChan, err := s.core.Subscribe(s.ctx, s.topic)
// 	if err != nil {
// 		s.logger.Bg().Error("Subcribe", zap.Error(err), zap.String("topic", s.topic))
// 		return err
// 	}

// 	s.msgChan = msgChan

// 	go s.process()

// 	return nil
// }

// func (s *Subscriber) process() {
// 	for {
// 		select {
// 		case msg := <-s.msgChan:

// 			if err := subscriber.ExecuteHandler(msg, s.logger); err != nil {
// 				s.logger.Bg().Error("Execute handle has error", zap.String("topic", s.topic), zap.Error(err))
// 			}

// 			msg.Ack()

// 		case <-s.ctx.Done():
// 			s.logger.Bg().Warn("Subscriber process stopped", zap.String("topic", s.topic))
// 			break
// 		}
// 	}
// }

// // Stop ...
// func (s *Subscriber) Stop() error {
// 	s.logger.Bg().Info("Transaction subscriber stopping", zap.String("topic", s.topic))
// 	s.cancelCtxFn()
// 	return nil
// }
