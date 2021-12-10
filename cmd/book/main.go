package main

import (
	"os"

	"github.com/blr-coder/book_grpc/internal/config"
	"github.com/blr-coder/book_grpc/internal/delivery/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	configPath = "/configs/config.toml"
)

func main() {
	app := &cli.App{
		Name:   "Book",
		Usage:  "gGRP server",
		Action: bookRun,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `FILE`",
				TakesFile:   true,
				Value:       configPath,
				DefaultText: configPath,
				Destination: &configPath,
				EnvVars:     []string{"BILLING_CONFIG_PATH"},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func bookRun(context *cli.Context) error {
	appConfig, err := initConfig(configPath)
	if err != nil {
		return err
	}
	logLevel, err := log.ParseLevel(appConfig.LogLevel)
	if err != nil {
		return err
	}
	logger := log.StandardLogger()
	log.SetLevel(logLevel)
	return grpc.RunServer(context.Context, appConfig, logger)
}

func initConfig(path string) (*config.Config, error) {
	appConfig, err := config.ParseConfig(path)
	if err != nil {
		return nil, err
	}
	return appConfig, nil
}
