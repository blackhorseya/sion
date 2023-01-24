// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/blackhorseya/irent/internal/adapter/account/restful"
	"github.com/blackhorseya/irent/internal/app/domain/account/biz"
	"github.com/blackhorseya/irent/internal/app/domain/account/biz/repo"
	"github.com/blackhorseya/irent/internal/pkg/config"
	"github.com/blackhorseya/irent/internal/pkg/httpx"
	"github.com/blackhorseya/irent/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateService(path2 string) (*Service, error) {
	viper, err := config.NewConfig(path2)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLogger(options)
	if err != nil {
		return nil, err
	}
	httpxOptions, err := httpx.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	engine := httpx.NewRouter(httpxOptions)
	server := httpx.NewServer(httpxOptions, logger, engine)
	repoOptions, err := repo.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client := httpx.NewClient()
	iRepo := repo.NewImpl(repoOptions, client)
	iBiz := biz.NewImpl(iRepo)
	adaptersRestful := restful.NewRestful(logger, engine, iBiz)
	service, err := NewService(logger, server, adaptersRestful)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, httpx.ProviderClientSet, httpx.ProviderServerSet, restful.AccountSet, biz.AccountSet, repo.AccountSet, NewService)
