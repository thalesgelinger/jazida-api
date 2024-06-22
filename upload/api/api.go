package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"upload/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ApiHandler struct {
	bucket *s3.S3
}

func NewApiHandler(s3Svc *s3.S3) ApiHandler {
	return ApiHandler{
		bucket: s3Svc,
	}
}

func (a *ApiHandler) HandlePreSignedUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	key := fmt.Sprintf("%d.jpg", time.Now().Unix())
	req, _ := a.bucket.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(5 * time.Minute)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, err
	}

	response := types.PreSignedUrl{
		URL: urlStr,
		Key: key,
	}

	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}

func (a *ApiHandler) GetShowUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	key := request.PathParameters["key"]

	if key == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Missing key in path parameters"}`,
		}, nil
	}

	req, _ := a.bucket.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, err
	}

	response := types.PreSignedUrl{
		URL: urlStr,
	}

	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil

}
