package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OlegBabakov/pow-server/pkg/logger/zap"

	"github.com/OlegBabakov/pow-server/config"
	"github.com/OlegBabakov/pow-server/internal/client"
)

const (
	AppName       = "client"
	ErrConfigInit = "failed config initialization"
)

func main() {
	appCtx, _ := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	cfg, err := config.NewConfig(appCtx, config.ClientConfig{})
	if err != nil {
		log.Fatal(ErrConfigInit, err)
	}

	logger := zap.NewZapLogger(cfg.Logger)
	logger.InitLogger(AppName)

	cl := client.InitWithConfig(cfg, logger)

	if err = cl.Start(appCtx, cfg.RequestCount); err != nil {
		logger.Fatal(err)
	}
}
