//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/irent/internal/adapter/account/restful"
	"github.com/blackhorseya/irent/internal/app/domain/account/biz"
	"github.com/blackhorseya/irent/internal/pkg/config"
	"github.com/blackhorseya/irent/internal/pkg/httpx"
	"github.com/blackhorseya/irent/internal/pkg/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,
	httpx.ProviderClientSet,

	// server
	httpx.ProviderServerSet,

	// domain
	restful.AccountSet,
	biz.ProviderSet,

	// main
	NewService,
)

func CreateService(path string) (*Service, error) {
	panic(wire.Build(providerSet))
}