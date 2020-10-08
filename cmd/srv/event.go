package main

import (
	"strings"

	"github.com/richard-xtek/go-grpc-micro-kit/kafka"
	"github.com/richard-xtek/go-grpc-micro-kit/subscriber"

	"xtek/exchange/commission/internal/delivery/msgevent"
	"xtek/exchange/commission/internal/events"
	pbLendingTypes "xtek/exchange/commission/pb/lending/types"

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
	eventDelivery := msgevent.New(s.comDomain)

	registry := subscriber.GetHandleRegistry()

	registry.Register(events.ETTSaveCommissionHistory, eventDelivery.CommissionEvent, &pbLendingTypes.CommissionEvent{})

	subscriberCfg := kafka.SubscriberConfig{
		Brokers:       strings.Split(s.cfg.KafkaBrokers, ","),
		Unmarshaler:   kafka.DefaultMarshaler{},
		ConsumerGroup: "commission_service",
	}

	kafkaSubscriber, err := kafka.NewSubscriber(subscriberCfg, s.logFactory)
	if err != nil {
		return err
	}

	CommissionSubscriber := events.NewSubscriber(s.logFactory.With(zap.String("subscriber_name", "commission.response")), kafkaSubscriber, events.CommissionTopic)
	subscriberWorkers.Register(CommissionSubscriber)

	s.registerProcessor(subscriberWorkers)
	return nil
}
