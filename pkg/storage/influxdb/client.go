package influxdb

import (
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/sion/pkg/configx"
)

// NewClient is a function to create a new influxdb client.
func NewClient() (*influxdb3.Client, error) {
	return influxdb3.New(influxdb3.ClientConfig{
		Host:  configx.C.Storage.Influxdb.URL,
		Token: configx.C.Storage.Influxdb.Token,
	})
}
