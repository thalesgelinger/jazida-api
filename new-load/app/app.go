package app

import (
	"new-load/api"
	db "new-load/database"
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
