package main

import (
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Username string `json:"username"`
}

func main() {
	lambdaApp := app.NewApp()
	lambda.Start(lambdaApp.ApiHandler.RegisterUser)
}
