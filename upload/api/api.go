package api

import (
	"context"
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

func (a *ApiHandler) HandlePreSignedUrl(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	key := fmt.Sprintf("uploads/%d.jpg", time.Now().Unix())
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
		UploadURL: urlStr,
		Key:       key,
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
