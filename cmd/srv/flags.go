package main

import (
	"os"
	"path/filepath"

	"xtek/exchange/commission/internal/config"

	"github.com/urfave/cli"
)

// Define flags
var (
	GRPCPortFlag = cli.StringFlag{
		Name:   "grpc.port",
		Usage:  "Port listen",
		EnvVar: "USER_GRPC_PORT",
		Value:  "10020",
	}
	GRPCHostFlag = cli.StringFlag{
		Name:   "grpc.host",
		Usage:  "Port listen",
		EnvVar: "USER_GRPC_HOST",
		Value:  "",
	}
	HTTPPortFlag = cli.StringFlag{
		Name:   "http.port",
		Usage:  "Port listen",
		EnvVar: "USER_HTTP_PORT",
		Value:  "11013",
	}
	HTTPHostFlag = cli.StringFlag{
		Name:   "http.host",
		Usage:  "HTTP Port to listen",
		EnvVar: "USER_HTTP_HOST",
		Value:  "",
	}

	MongoHostFlag = cli.StringFlag{
		Name:   "mongo.host",
		Usage:  "Mongo host used connect mongoDB",
		Value:  "localhost",
		EnvVar: "MONGO_HOST",
	}
	MongoPortFlag = cli.StringFlag{
		Name:   "mongo.port",
		Usage:  "Mongo port used connect mongoDB",
		Value:  "27017",
		EnvVar: "MONGO_PORT",
	}
	MongoDataBaseFlag = cli.StringFlag{
		Name:   "mongo.db",
		Usage:  "Mongo port used connect mongoDB",
		Value:  "Wallet",
		EnvVar: "MONGO_DB",
	}
	MongoUsernameFlag = cli.StringFlag{
		Name:   "mongo.username",
		Usage:  "Mongo port used connect mongoDB",
		Value:  "",
		EnvVar: "MONGO_USERNAME",
	}
	MongoPasswordFlag = cli.StringFlag{
		Name:   "mongo.password",
		Usage:  "Mongo password used connect mongoDB",
		Value:  "",
		EnvVar: "MONGO_PASSWORD",
	}

	ServiceNameFlag = cli.StringFlag{
		Name:   "name",
		Usage:  "Service name",
		EnvVar: "USER_SERVICE_NAME",
		Value:  "Wallet",
	}
	JaegerAddressFlag = cli.StringFlag{
		Name:   "jaeger_address",
		Usage:  "Jaeger Address used to connect",
		EnvVar: "USER_JAEGER_ADDRESS",
		Value:  "localhost:6831",
	}
	ConsulAddressFlag = cli.StringFlag{
		Name:   "consul_address",
		Usage:  "Consul Address used to connect",
		EnvVar: "USER_CONSUL_ADDRESS",
		Value:  "localhost:8500",
	}
	PrometheusPortFlag = cli.StringFlag{
		Name:   "prometheus.port",
		Usage:  "Prometheus port",
		EnvVar: "PROMETHEUS_PORT",
		Value:  "9010",
	}
	CoinbaseServiceNameFlag = cli.StringFlag{
		Name:   "services.coinbase_name",
		Usage:  "services name of coinbase services",
		EnvVar: "SERVICES_COINBASE_NAME",
		Value:  "Coinbase",
	}
	CryptocurrencyServiceNameFlag = cli.StringFlag{
		Name:   "services.cryptocurrency_name",
		Usage:  "services name of cryptocurrency services",
		EnvVar: "SERVICES_CRYPTOCURRENCY_NAME",
		Value:  "Cryptocurrency",
	}
	OrderServiceNameFlag = cli.StringFlag{
		Name:   "services.order_name",
		Usage:  "services name of order services",
		EnvVar: "SERVICES_ORDER_NAME",
		Value:  "Order",
	}
	UserServiceNameFlag = cli.StringFlag{
		Name:   "services.user_name",
		Usage:  "services name of user services",
		EnvVar: "SERVICES_USER_NAME",
		Value:  "User",
	}
	VerifyCodeServiceNameFlag = cli.StringFlag{
		Name:   "services.verify_code_name",
		Usage:  "services name of verify code services",
		EnvVar: "SERVICES_VERIFYCODE_NAME",
		Value:  "VerifyCode",
	}
	LendingServiceNameFlag = cli.StringFlag{
		Name:   "services.lending_name",
		Usage:  "services name of lending services",
		EnvVar: "SERVICES_LENDING_NAME",
		Value:  "Lending",
	}
	KafkaBrokersFlag = cli.StringFlag{
		Name:   "kafka.brokers",
		Usage:  "Kafka brokers",
		EnvVar: "KAFKA_BROKERS",
		Value:  "localhost:9092",
	}
)

var (
	startCommand = cli.Command{
		Action:      migrateFlags(s.start),
		Name:        "start",
		Usage:       "Bootstrap and start worker server",
		ArgsUsage:   "<genesisPath>",
		Flags:       []cli.Flag{},
		Category:    "Crawler Worker",
		Description: `Used to start crawler worker, clone data from omada cloud`,
	}
	eventCommand = cli.Command{
		Action:      migrateFlags(s.startEventProcess),
		Name:        "event",
		Usage:       "Bootstrap and start worker server",
		ArgsUsage:   "<genesisPath>",
		Flags:       []cli.Flag{},
		Category:    "Crawler Worker",
		Description: `Used to start crawler worker, clone data from omada cloud`,
	}
)

// NewApp creates an app with sane defaults.
func NewApp(gitCommit, gitDate, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "Richard Pham"
	app.Email = "richard@xtek.asia"
	app.Version = VersionWithCommit(gitCommit, gitDate)
	app.Usage = usage
	return app
}

// makeConfig ...
func makeConfig(ctx *cli.Context) *config.Config {
	return &config.Config{
		HTTP: config.ConnAddress{
			Host: ctx.GlobalString(HTTPHostFlag.GetName()),
			Port: ctx.GlobalString(HTTPPortFlag.GetName()),
		},
		GRPC: config.ConnAddress{
			Host: ctx.GlobalString(GRPCHostFlag.GetName()),
			Port: ctx.GlobalString(GRPCPortFlag.GetName()),
		},
		Mongo: config.Mongo{
			Address:  ctx.GlobalString(MongoHostFlag.GetName()),
			Port:     ctx.GlobalString(MongoPortFlag.GetName()),
			Database: ctx.GlobalString(MongoDataBaseFlag.GetName()),
			Username: ctx.GlobalString(MongoUsernameFlag.GetName()),
			Password: ctx.GlobalString(MongoPasswordFlag.GetName()),
		},
		ServiceName:   ctx.GlobalString(ServiceNameFlag.GetName()),
		ConsulAddress: ctx.GlobalString(ConsulAddressFlag.GetName()),
		JaegerAddress: ctx.GlobalString(JaegerAddressFlag.GetName()),

		PrometheusPort: ctx.GlobalString(PrometheusPortFlag.GetName()),

		CoinbaseService:       ctx.GlobalString(CoinbaseServiceNameFlag.GetName()),
		CryptocurrencyService: ctx.GlobalString(CryptocurrencyServiceNameFlag.GetName()),
		OrderService:          ctx.GlobalString(OrderServiceNameFlag.GetName()),
		UserService:           ctx.GlobalString(UserServiceNameFlag.GetName()),
		VerifyCodeService:     ctx.GlobalString(VerifyCodeServiceNameFlag.GetName()),
		LendingService:        ctx.GlobalString(LendingServiceNameFlag.GetName()),

		KafkaBrokers: ctx.GlobalString(KafkaBrokersFlag.GetName()),
	}
}
