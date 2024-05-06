package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Username string `json:"username"`
}

func HandleRequest(payload Payload) (string, error) {
	if payload.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by %s", payload.Username), nil

}
func main() {
	lambda.Start(HandleRequest)
}
