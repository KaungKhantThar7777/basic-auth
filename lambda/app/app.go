package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	db := database.NewDynamoDB()
	api := api.NewApiHandler(db)

	return App{
		ApiHandler: api,
	}
}
