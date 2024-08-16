// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/sion/app/domain/rental/repo/asset/irent"
	"github.com/blackhorseya/sion/entity/domain/rental/repo"
	"github.com/blackhorseya/sion/pkg/storage/influxdbx"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, error) {
	iAssetRepo, err := irent.NewAssetRepo()
	if err != nil {
		return nil, err
	}
	client, err := influxdbx.NewClient()
	if err != nil {
		return nil, err
	}
	mainInjector := &Injector{
		assets:   iAssetRepo,
		influxdb: client,
	}
	return mainInjector, nil
}

// wire.go:

// Injector is the wire injector for the fetchAvailableCars job.
type Injector struct {
	assets   repo.IAssetRepo
	influxdb *influxdb3.Client
}