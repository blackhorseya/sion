//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/irent/internal/pkg/config"
	"github.com/blackhorseya/irent/internal/pkg/log"
	"github.com/blackhorseya/irent/pkg/app"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,

	NewService,
)

func CreateApplication(path string) (app.Servicer, error) {
	panic(wire.Build(providerSet))
}
