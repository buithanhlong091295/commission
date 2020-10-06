package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/richard-xtek/go-grpc-micro-kit/log"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"

	"github.com/richard-xtek/go-grpc-micro-kit/tracing"

	grpcDelivery "xtek/exchange/commission/internal/delivery/grpc"

	"github.com/opentracing/opentracing-go"

	pbCom "xtek/exchange/commission/pb/commission"

	"github.com/richard-xtek/go-grpc-micro-kit/server"

	"xtek/exchange/commission/internal/config"

	consul "github.com/hashicorp/consul/api"
	"github.com/richard-xtek/go-grpc-micro-kit/prometheus"
	"github.com/richard-xtek/go-grpc-micro-kit/registry"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

// ConnectMongo
func connectMongo(cfg *config.Mongo) (*mgo.Session, error) {
	s := cfg.ConnectionString()
	session, err := mgo.Dial(s)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Processor processes metrics in multiple formats
type Processor interface {
	Start() error
	Stop() error
}

type srv struct {
	cfg *config.Config

	logFactory log.Factory

	// define for stores

	stopChan chan struct{}

	processors []Processor

	prometheusFactory metrics.Factory
	tracer            opentracing.Tracer
	consul            *consul.Client

	mgoSession *mgo.Session

	// coinBaseClient   pbCoin.InternalSiteServiceClient
	// cryptoClient     pbCrypto.InternalSiteServiceClient
	// orderClient      pbOrder.InternalSiteServiceClient
	// userClient       pbUser.InternalSiteServiceClient
	// verifyCodeClient pbVerifyCode.InternalSiteServiceClient
	// lendingClient    pbLending.InternalSiteServiceClient

	// publisher *events.Publisher
}

func (s *srv) start(ctx *cli.Context) error {
	if err := s.prepare(ctx); err != nil {
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

func (s *srv) makeLogger() {
	var logger log.Factory
	if s.cfg.IsProduction && s.cfg.LogFolderPath != "" {
		logger = log.NewStandardFactory(s.cfg.LogFolderPath, s.cfg.ServiceName)
	} else {
		logger = log.NewDevelopFactory(s.cfg.ServiceName)
	}
	s.logFactory = logger
}

func (s *srv) prepare(ctx *cli.Context) (err error) {
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

	// if err := s.loadPublisher(); err != nil {
	// 	return err
	// }

	if err := s.loadStores(); err != nil {
		return err
	}

	if err := s.loadGRPCClient(); err != nil {
		return err
	}

	if err := s.loadDomains(); err != nil {
		return err
	}

	if err := s.loadCrons(); err != nil {
		return err
	}

	if err := s.loadGRPCServer(); err != nil {
		return err
	}

	return nil
}

func (s *srv) loadPublisher() error {
	// publishCfg := kafka.PublisherConfig{
	// 	Brokers:   strings.Split(s.cfg.KafkaBrokers, ","),
	// 	Marshaler: kafka.DefaultMarshaler{},
	// }

	// publisher, err := kafka.NewPublisher(publishCfg, s.logFactory)
	// if err != nil {
	// 	return err
	// }
	// s.publisher = events.NewPublisher(publisher)
	return nil
}

func (s *srv) loadStores() (err error) {
	// s.balanceChangeLogStore, err = mongo.NewBalanceChangeLogRepo(s.cfg.Mongo.Database, s.mgoSession)
	// if err != nil {
	// 	return err
	// }

	return
}

func (s *srv) loadCrons() error {
	// cron1 := cron.NewAutoApprovalWithdrawCron(s.logFactory, s.withdrawDomain)
	// s.registerProcessor(cron1)

	return nil
}

func (s *srv) loadGRPCClient() error {
	// clientDialer := grpc_dialer.NewGrpcClientDialer(s.consul, s.tracer, s.logFactory)

	// lendingConn, err := clientDialer.ConnWithServiceName(s.cfg.LendingService)
	// if err != nil {
	// 	return err
	// }

	// s.lendingClient = pbLending.NewInternalSiteServiceClient(lendingConn)
	return nil
}

func (s *srv) loadDomains() error {
	// s.transferDomain = dTransfer.New(s.balanceStore, s.coinStore, s.transactionStore, s.transferHistoryStore, s.balanceChangeLogStore, s.orderClient, s.lendingClient, s.publisher)

	return nil
}

func (s *srv) loadGRPCServer() error {
	consulRegistrar := registry.NewConsulRegister(
		s.consul,
		s.cfg.ServiceName,
		s.cfg.GRPC.GetPortInt(),
		[]string{"wallet"},
	)

	userSite, err := grpcDelivery.NewUserSiteDelivery()
	if err != nil {
		return err
	}

	// healthCheck := grpcDelivery.NewHealthService()

	grpcServer := server.NewGRPCServer(s.logFactory, s.cfg.ServiceName).WithHandler(func(s *grpc.Server) {
		pbCom.RegisterUserSiteServiceServer(s, userSite)
		// pbWallet.RegisterAdminSiteServiceServer(s, adminSite)
		// pbWallet.RegisterInternalSiteServiceServer(s, internalSite)
		// pbHealth.RegisterHealthServer(s, healthCheck)
	}).WithHost(s.cfg.GRPC.Host).WithPort(s.cfg.GRPC.Port).WithConsul(consulRegistrar).WithTracer(s.tracer)

	if s.cfg.PrometheusPort != "" {
		grpcServer.EnablePrometheus(s.cfg.PrometheusPort, true)
	}

	// register to processors
	s.registerProcessor(grpcServer)
	return nil
}

func (s *srv) signer() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	<-sigc
	s.logFactory.Bg().Info("Got interrupt, shutting down...")
	go s.stop()

	for i := 10; i > 0; i-- {
		<-sigc
		if i > 1 {
			s.logFactory.Bg().Warn("Already shutting down, interrupt more to panic.", zap.Int("times", i-1))
		}
	}
}

func (s *srv) registerProcessor(processor Processor) {
	s.processors = append(s.processors, processor)
}

func (s *srv) stop() error {
	// unblock n.Wait
	close(s.stopChan)
	return nil
}

func (s *srv) wait() {
	<-s.stopChan
}
