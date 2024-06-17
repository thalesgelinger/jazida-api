package api

import (
	"fmt"
	db "new-load/database"
	"new-load/types"

	"github.com/google/uuid"
)

type ApiHandler struct {
	db db.DynamoDbClient
}

func NewApiHandler(db db.DynamoDbClient) ApiHandler {
	return ApiHandler{
		db: db,
	}
}

func (a *ApiHandler) NewLoadHandler(load types.Load) error {

	if load.Client == "" {
		return fmt.Errorf("Missing client field")
	}

	if load.Plate == "" {
		return fmt.Errorf("Missing plate field")
	}

	if load.Material == "" {
		return fmt.Errorf("Missing material field")
	}

	if load.Quantity == "" {
		return fmt.Errorf("Missing quantity field")
	}

	if load.PaymentMethod == "" {
		return fmt.Errorf("Missing payment method field")
	}

	if load.Signature == "" {
		return fmt.Errorf("Missing signature field")
	}

	load.Id = uuid.New().String()

	err := a.db.InsertLoad(load)
	if err != nil {
		return fmt.Errorf("Error inserting Load: %w", err)
	}
	return nil
}
