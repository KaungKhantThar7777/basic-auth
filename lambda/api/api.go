package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamodbClient
}

func NewApiHandler(dbStore database.DynamodbClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api *ApiHandler) RegisterUser(payload types.RegisterUser) error {
	if payload.Username == "" || payload.Password == "" {
		return fmt.Errorf("invalid request, the fields cannot be empty")
	}

	isExisted, error := api.dbStore.DoesUserExists(payload.Username)

	if error != nil {
		return fmt.Errorf("there was a problem %w", error)
	}

	if isExisted {
		return fmt.Errorf("a user with that username already existed")
	}

	error = api.dbStore.InsertUser(payload)

	if error != nil {
		return fmt.Errorf("error inserting the user %w", error)
	}

	return nil
}
