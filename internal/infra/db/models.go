// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

type Client struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Load struct {
	ID            int64       `json:"id"`
	ClientID      int64       `json:"client_id"`
	PlateID       int64       `json:"plate_id"`
	MaterialID    int64       `json:"material_id"`
	Quantity      string      `json:"quantity"`
	PaymentMethod string      `json:"payment_method"`
	Signature     string      `json:"signature"`
	Foreign       interface{} `json:"foreign"`
}

type Material struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Plate struct {
	ID       int64  `json:"id"`
	ClientID int64  `json:"client_id"`
	Plate    string `json:"plate"`
}
