package main

import (
	"github.com/richard-xtek/go-grpc-micro-kit/subscriber"

	"github.com/richard-xtek/go-grpc-micro-kit/prometheus"
	"github.com/richard-xtek/go-grpc-micro-kit/registry"
	"github.com/richard-xtek/go-grpc-micro-kit/tracing"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func (s *srv) startEventProcess(ctx *cli.Context) error {
	if err := s.prepareEventProcess(ctx); err != nil {
		return err
	}

	for _, processor := range s.processors {
		if err := processor.Start(); err != nil {
			return err
		}
	}

	go s.signer()

	s.wait()
	return nil
}

func (s *srv) prepareEventProcess(ctx *cli.Context) (err error) {
	s.cfg = makeConfig(ctx)
	s.makeLogger()

	s.stopChan = make(chan struct{})

	prometheusBuilder := prometheus.New("")
	s.prometheusFactory, err = prometheusBuilder.CreateMetricsFactory("")
	if err != nil {
		return err
	}

	s.tracer = tracing.Init(s.cfg.ServiceName, s.logFactory, s.cfg.JaegerAddress)
	// Create Consul
	s.consul, err = registry.NewClient(s.cfg.ConsulAddress)
	if err != nil {
		s.logFactory.Bg().Fatal("Consul register failure", zap.Error(err))
	}

	s.mgoSession, err = connectMongo(&s.cfg.Mongo)
	if err != nil {
		s.logFactory.Bg().Fatal("Mongo register failure", zap.Error(err))
	}

	if err := s.loadPublisher(); err != nil {
		return err
	}

	if err := s.loadStores(); err != nil {
		return err
	}

	if err := s.loadGRPCClient(); err != nil {
		return err
	}

	if err := s.loadDomains(); err != nil {
		return err
	}

	if err := s.loadKafkaSubscriber(); err != nil {
		return err
	}

	return nil
}

func (s *srv) loadKafkaSubscriber() error {
	subscriberWorkers := subscriber.GetSubscriberWorker()
	// eventDelivery := msgevent.New()

	// registry := subscriber.GetHandleRegistry()

	// registry.Register(events.ETTOrderPlaceResponse, eventDelivery.OrderMatchingEngineResponse, &pbOrderTypes.OrderMatchingEngineResponse{})
	// registry.Register(events.ETTTransactionBitcoinUnConfirmed, eventDelivery.TransactionBitcoinUnconfirmed, &pbWalletTypes.TransactionBitcoin{})
	// registry.Register(events.ETTTransactionBitcoinConfirmed, eventDelivery.TransactionBitcoinConfirmed, &pbWalletTypes.TransactionBitcoin{})
	// registry.Register(events.ETTTransactionWithdrawConfirmed, eventDelivery.TransactionWithdrawConfirmed, &pbWalletTypes.TransactionInternal{})
	// registry.Register(events.ETTAdminApprovalWithdraw, eventDelivery.AdminApprovalWithdraw, &pbWalletTypes.TransactionInternal{})
	// registry.Register(events.ETTTransactionDepositPending, eventDelivery.TransactionDepositPending, &pbWalletTypes.TransactionDepositEvent{})
	// registry.Register(events.ETTTransactionDepositCompleted, eventDelivery.TransactionDepositCompleted, &pbWalletTypes.TransactionDepositEvent{})
	// registry.Register(events.ETTTransferPending, eventDelivery.HandleTransferPending, &pbWalletTypes.Transfer{})
	// registry.Register(events.ETTContractApplyCheckoutRequestEvent, eventDelivery.LoanContractApplyCheckout, &pbLendingTypes.Contract{})

	// subscriberCfg := kafka.SubscriberConfig{
	// 	Brokers:       strings.Split(s.cfg.KafkaBrokers, ","),
	// 	Unmarshaler:   kafka.DefaultMarshaler{},
	// 	ConsumerGroup: "wallet_service",
	// }

	// kafkaSubscriber, err := kafka.NewSubscriber(subscriberCfg, s.logFactory)
	// if err != nil {
	// 	return err
	// }

	// walletEngineSubscriber := events.NewSubscriber(s.logFactory.With(zap.String("subscriber_name", "wallet_engine")), kafkaSubscriber, events.WalletEngineTopic)
	// subscriberWorkers.Register(walletEngineSubscriber)

	// transactionPendingSubscriber := events.NewSubscriber(s.logFactory.With(zap.String("subscriber_name", "transaction_pending")), kafkaSubscriber, events.Topic.TransactionPending())
	// subscriberWorkers.Register(transactionPendingSubscriber)

	// withdrawApprovalEngineSubscriber := events.NewSubscriber(s.logFactory.With(zap.String("subscriber_name", "withdraw_approval_engine")), kafkaSubscriber, events.WithdrawApprovalEngineTopic)
	// subscriberWorkers.Register(withdrawApprovalEngineSubscriber)

	s.registerProcessor(subscriberWorkers)
	return nil
}
