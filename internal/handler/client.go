package handler

import (
	"encoding/json"
	"jazida-api/internal/infra/db"
	"log"
	"net/http"
)

type ClientHandler struct {
	db *db.Queries
}

func NewClientHandler(db *db.Queries) *ClientHandler {
	return &ClientHandler{
		db: db,
	}
}

type ClientsResponse struct {
	Name   string   `json:"name"`
	Plates []string `json:"plates"`
}

func (c *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := c.db.GetClients(r.Context())
	if err != nil {
		log.Fatal("Error getting clients")
		return
	}

	clientsResponse := []ClientsResponse{}

	responseMap := map[string][]string{}

	for _, client := range clients {
		if len(responseMap[client.Name]) == 0 {
			responseMap[client.Name] = []string{client.Plate}
		} else {
			responseMap[client.Name] = append(responseMap[client.Name], client.Plate)
		}
	}

	for name, plates := range responseMap {
		newClientResponse := ClientsResponse{
			Name:   name,
			Plates: plates,
		}
		clientsResponse = append(clientsResponse, newClientResponse)
	}

	err = json.NewEncoder(w).Encode(clientsResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
