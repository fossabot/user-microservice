package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thomaspoignant/user-microservice/config"
	"github.com/thomaspoignant/user-microservice/server"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func main() {
	//load config file
	config.LoadConfigFile()
	switch runAs := viper.GetString("RUNNING_MODE"); runAs {
	case "lambda":
		log.Info("Run as lambda")
		lambdaRun()
	default:
		log.Info("Run locally")
		localRun()
	}
}

func localRun() {
	server.Init()
}

func lambdaRun() {
	lambda.Start(Handler)
}

// Handler to wrap gin to lambda
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		router := server.SetupRouter()
		ginLambda = ginadapter.New(router)
		initialized = true
	}

	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.Proxy(req)
}
