// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Client struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Load struct {
	ID            int32            `json:"id"`
	ClientID      int32            `json:"client_id"`
	PlateID       int32            `json:"plate_id"`
	MaterialID    int32            `json:"material_id"`
	Quantity      string           `json:"quantity"`
	PaymentMethod string           `json:"payment_method"`
	Signature     string           `json:"signature"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
}

type Material struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Plate struct {
	ID       int32  `json:"id"`
	ClientID int32  `json:"client_id"`
	Plate    string `json:"plate"`
}
