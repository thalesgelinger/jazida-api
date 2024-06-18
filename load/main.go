package main

import (
	"load/app"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := app.NewApp()
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.Path {
		case "/loads":
			switch req.HTTPMethod {
			case "POST":
				return app.ApiHandler.NewLoadHandler(req)
			case "GET":
				return app.ApiHandler.GetAllLoads(req)
			default:
				return events.APIGatewayProxyResponse{
					Body:       "Not found",
					StatusCode: http.StatusNotFound,
				}, nil

			}
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
