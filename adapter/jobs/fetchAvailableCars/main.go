package main

import (
	"net/http"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/sion/pkg/configx"
	"github.com/blackhorseya/sion/pkg/contextx"
	"github.com/blackhorseya/sion/pkg/logging"
	"go.uber.org/zap"
)

const (
	database = "irent_cars"
)

var injector *Injector

// Handler is our lambda handler invoked by the `lambda.Start` function call.
func Handler() (events.APIGatewayProxyResponse, error) {
	ctx := contextx.Background()
	cars, err := injector.assets.FetchAvailableCars(ctx)
	if err != nil {
		ctx.Error("failed to fetch available cars", zap.Error(err))
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	now := time.Now()
	var points []*influxdb3.Point
	for _, car := range cars {
		points = append(points, influxdb3.NewPointWithMeasurement("available_cars").
			SetTag("car_id", car.ID).
			SetDoubleField("latitude", car.Location.Latitude).
			SetDoubleField("longitude", car.Location.Longitude).
			SetTimestamp(now))
	}

	err = injector.influxdb.WritePoints(ctx, points, influxdb3.WithDatabase(database))
	if err != nil {
		ctx.Error("failed to write points", zap.Error(err))
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	err := configx.LoadWithPathAndName("", "sion")
	if err != nil {
		panic(err)
	}

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		panic(err)
	}

	injector, err = BuildInjector()
	if err != nil {
		panic(err)
	}

	lambda.Start(Handler)
}
