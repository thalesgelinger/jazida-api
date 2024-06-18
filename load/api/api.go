package api

import (
	"encoding/json"
	db "load/database"
	"load/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

type ApiHandler struct {
	db db.DBClient
}

func NewApiHandler(db db.DBClient) ApiHandler {
	return ApiHandler{
		db: db,
	}
}

func (a *ApiHandler) NewLoadHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var load types.Load

	err := json.Unmarshal([]byte(req.Body), load)
	if err != nil {
		return writeError(http.StatusBadRequest, "Invalid request"), err
	}

	if load.Client == "" {
		return writeError(http.StatusBadRequest, "Missing client field"), err
	}

	if load.Plate == "" {
		return writeError(http.StatusBadRequest, "Missing plate field"), err
	}

	if load.Material == "" {
		return writeError(http.StatusBadRequest, "Missing material field"), err
	}

	if load.Quantity == "" {
		return writeError(http.StatusBadRequest, "Missing quantity field"), err
	}

	if load.PaymentMethod == "" {
		return writeError(http.StatusBadRequest, "Missing payment method field"), err
	}

	if load.Signature == "" {
		return writeError(http.StatusBadRequest, "Missing signature field"), err
	}

	load.Id = uuid.New().String()

	err = a.db.InsertLoad(load)
	if err != nil {
		return writeError(http.StatusInternalServerError, "Error inserting load"), err
	}
	return events.APIGatewayProxyResponse{}, nil
}

func (a *ApiHandler) GetAllLoads(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	loads, err := a.db.GetAllLoads()

	if err != nil {
		return writeError(http.StatusInternalServerError, "Error getting all loads"), err
	}

	loadsJSON, err := json.Marshal(loads)

	if err != nil {
		return writeError(http.StatusInternalServerError, "Error parsing loads to json"), err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(loadsJSON),
		StatusCode: http.StatusOK,
	}, nil
}

func writeError(status int, msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: status,
	}
}
