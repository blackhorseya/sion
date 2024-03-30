package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blackhorseya/sion/pkg/configx"
	"github.com/blackhorseya/sion/pkg/contextx"
	"github.com/blackhorseya/sion/pkg/logging"
	"go.uber.org/zap"
)

var injector *Injector

// Handler is our lambda handler invoked by the `lambda.Start` function call.
func Handler() (events.APIGatewayProxyResponse, error) {
	ctx := contextx.Background()
	cars, err := injector.assets.FetchAvailableCars(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	ctx.Info("fetched available cars", zap.Any("cars", cars[:5]))

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
