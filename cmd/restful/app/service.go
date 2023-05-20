package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/irent/pkg/app"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
}

// NewService creates a new app.Servicer instance
func NewService(logger *zap.Logger) app.Servicer {
	return &impl{
		logger: logger,
	}
}

func (i *impl) Start() error {
	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		i.logger.Info("receive a signal", zap.String("signal", sig.String()))

		os.Exit(0)
	}

	return nil
}
