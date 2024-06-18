package app

import (
	"load/api"
	db "load/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {

	db := db.NewDynamoDb()
	api := api.NewApiHandler(db)

	return App{
		ApiHandler: api,
	}
}
