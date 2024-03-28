package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/sion/adapter/restful"
	"github.com/blackhorseya/sion/pkg/configx"
	"github.com/blackhorseya/sion/pkg/logging"
)

var ginLambda *ginadapter.GinLambda

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	err := configx.LoadWithPathAndName("", "sion")
	if err != nil {
		log.Fatal(err)
	}

	err = logging.InitWithConfig(configx.C.Log)
	if err != nil {
		log.Fatal(err)
	}

	app, err := restful.NewRestful()
	if err != nil {
		log.Fatal(err)
	}

	err = app.InitRouting()
	if err != nil {
		log.Fatal(err)
	}

	ginLambda = ginadapter.New(app.GetRouter())

	lambda.Start(Handler)
}
