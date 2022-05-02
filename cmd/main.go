package main

import (
	"context"
	"eduid-navet/internal/apiv2"
	"eduid-navet/internal/httpserver"
	"eduid-navet/pkg/configuration"
	"eduid-navet/pkg/logger"
	"eduid-navet/pkg/model"
	"os"
	"os/signal"
	"syscall"
)

type service interface {
	Close(ctx context.Context) error
}

func main() {
	ctx := context.Background()
	cfg := &model.Cfg{}

	var (
		log      *logger.Logger
		mainLog  *logger.Logger
		services = make(map[string]service)
	)

	cfg, err := configuration.Parse(logger.NewSimple("Configuration"))
	if err != nil {
		panic(err)
	}

	mainLog = logger.New("main", cfg.Production)
	log = logger.New("navet", cfg.Production)

	apiv2, err := apiv2.New(ctx, cfg, log.New("apiv1"))
	if err != nil {
		panic(err)
	}

	httpserver, err := httpserver.New(ctx, cfg, apiv2, log.New("httpserver"))
	if err != nil {
		panic(err)
	}
	services["httpserver"] = httpserver

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	mainLog.Info("HALTING SIGNAL!")

	for serviceName, service := range services {
		if err := service.Close(ctx); err != nil {
			mainLog.Warn(serviceName)
		}
	}

	mainLog.Info("Stopped")
}
