package main

import (
	"new-load/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := app.NewApp()
	lambda.Start(app.ApiHandler.NewLoadHandler)
}
