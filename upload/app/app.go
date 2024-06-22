package app

import (
	"upload/api"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {

	sess := session.Must(session.NewSession())
	s3Svc := s3.New(sess)
	apiHandler := api.NewApiHandler(s3Svc)

	return App{
		ApiHandler: apiHandler,
	}
}
