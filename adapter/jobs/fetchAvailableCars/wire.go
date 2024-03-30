//go:build wireinject

//go:generate wire

package main

import (
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/sion/app/domain/rental/repo/asset/irent"
	"github.com/blackhorseya/sion/entity/domain/rental/repo"
	"github.com/blackhorseya/sion/pkg/storage/influxdbx"
	"github.com/google/wire"
)

// Injector is the wire injector for the fetchAvailableCars job.
type Injector struct {
	assets   repo.IAssetRepo
	influxdb *influxdb3.Client
}

func BuildInjector() (*Injector, error) {
	panic(wire.Build(
		wire.Struct(new(Injector), "*"),
		irent.NewAssetRepo,
		influxdbx.NewClient,
	))
}
