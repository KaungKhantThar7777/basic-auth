package api

import (
	"encoding/json"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api *ApiHandler) RegisterUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &payload)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request",
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	if payload.Username == "" || payload.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "invalid request, the fields cannot be empty",
			StatusCode: http.StatusBadRequest,
		}, fmt.Errorf("register user fields cannot be empty")
	}

	isExisted, error := api.dbStore.DoesUserExists(payload.Username)

	if error != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal error")
	}

	if isExisted {
		return events.APIGatewayProxyResponse{
			Body:       "Already existed",
			StatusCode: http.StatusConflict,
		}, fmt.Errorf("a user with that username already existed")

	}

	user, err := types.NewUser(payload)

	error = api.dbStore.InsertUser(user)

	if error != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal error")
	}

	return events.APIGatewayProxyResponse{
		Body:       "success",
		StatusCode: http.StatusOK,
	}, nil
}

func (api *ApiHandler) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var payload types.LoginUser

	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	user, err := api.dbStore.GetUser(payload.Username)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal srever",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if !types.ValidatePassword(user.PasswordHash, payload.Password) {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid login credentials",
			StatusCode: http.StatusUnauthorized,
		}, fmt.Errorf("")
	}

	return events.APIGatewayProxyResponse{
		Body:       "Logged in.",
		StatusCode: http.StatusOK,
	}, nil
}
