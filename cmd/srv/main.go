package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	gitDate   = ""
	// The app that holds all commands and flags.
	app = NewApp(gitCommit, gitDate, "the owifi crawler worker command line interface")

	s = new(srv)
)

func init() {
	app.Action = cli.ShowAppHelp
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2019-2020 The Owifi Crawler Worker Authors"
	app.Commands = []cli.Command{
		startCommand,
		eventCommand,
	}
	app.Flags = []cli.Flag{
		GRPCHostFlag,
		GRPCPortFlag,
		HTTPHostFlag,
		HTTPPortFlag,
		ServiceNameFlag,

		MongoDataBaseFlag,
		MongoHostFlag,
		MongoPortFlag,

		PrometheusPortFlag,

		CoinbaseServiceNameFlag,
		CryptocurrencyServiceNameFlag,
		OrderServiceNameFlag,
		UserServiceNameFlag,
		VerifyCodeServiceNameFlag,
		LendingServiceNameFlag,

		JaegerAddressFlag,
		ConsulAddressFlag,
		KafkaBrokersFlag,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
