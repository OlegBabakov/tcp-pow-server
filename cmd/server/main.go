package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OlegBabakov/pow-server/internal/server"
	"github.com/OlegBabakov/pow-server/pkg/logger/zap"

	"github.com/OlegBabakov/pow-server/config"
)

const (
	AppName       = "server"
	ErrConfigInit = "failed config initialization"
)

func main() {
	appCtx, _ := signal.NotifyContext(context.TODO(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	cfg, err := config.NewConfig(appCtx, config.ServerConfig{})
	if err != nil {
		log.Fatal(ErrConfigInit, err)
	}

	logger := zap.NewZapLogger(cfg.Logger)
	logger.InitLogger(AppName)

	srv := server.InitWithConfig(cfg, logger)

	go func() {
		<-appCtx.Done()
		logger.Info("Graceful shutdown...")
		srv.Stop()
	}()

	if err = srv.Run(appCtx); err != nil {
		logger.Fatal(err)
	}
}
