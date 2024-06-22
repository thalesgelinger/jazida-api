package main

import (
	"net/http"
	"upload/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := app.NewApp()
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.HTTPMethod {
		case "POST":
			return app.ApiHandler.HandlePreSignedUrl(req)
		case "GET":
			return app.ApiHandler.GetShowUrl(req)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
