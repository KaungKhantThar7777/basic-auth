package main

import (
	"lambda-func/app"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Username string `json:"username"`
}

func main() {
	lambdaApp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return lambdaApp.ApiHandler.RegisterUser(request)
		case "/login":
			return lambdaApp.ApiHandler.LoginUser(request)
		}
		return events.APIGatewayProxyResponse{
			Body:       "Not found",
			StatusCode: http.StatusNotFound,
		}, nil
	})
}
