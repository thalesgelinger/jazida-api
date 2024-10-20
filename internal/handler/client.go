package handler

import (
	"encoding/json"
	"jazida-api/internal/infra/db"
	"log"
	"net/http"
	"strconv"
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
	Id     int                         `json:"id"`
	Name   string                      `json:"name"`
	Plates []db.GetPlatesByClientIdRow `json:"plates"`
}

func (c *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := c.db.GetClients(r.Context())
	if err != nil {
		log.Fatal("Error getting clients")
		return
	}

	clientsResponse := []ClientsResponse{}

	for _, client := range clients {
		plates, err := c.db.GetPlatesByClientId(r.Context(), client.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clientsResponse = append(clientsResponse, ClientsResponse{
			Id:     int(client.ID),
			Name:   client.Name,
			Plates: plates,
		})
	}

	err = json.NewEncoder(w).Encode(clientsResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {

	type Client struct {
		Name string `json:"name"`
	}

	var client Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err := c.db.AddClient(r.Context(), client.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (c *ClientHandler) CreatePlate(w http.ResponseWriter, r *http.Request) {
	clientIdStr := r.PathValue("id")

	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type NewPlate struct {
		Plate string `json:"plate"`
	}

	var newPlate NewPlate
	if err := json.NewDecoder(r.Body).Decode(&newPlate); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err = c.db.AddPlate(r.Context(), db.AddPlateParams{
		ClientID: int64(clientId),
		Plate:    newPlate.Plate,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}

func (c *ClientHandler) GetMaterials(w http.ResponseWriter, r *http.Request) {
	materials, err := c.db.GetMaterials(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(materials)

	if err != nil {
		http.Error(w, "Error sending materials", http.StatusInternalServerError)
		return
	}

}
func (c *ClientHandler) CreateMaterial(w http.ResponseWriter, r *http.Request) {

	type NewMaterial struct {
		Material string `json:"material"`
	}

	var newMaterial NewMaterial

	if err := json.NewDecoder(r.Body).Decode(&newMaterial); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	err := c.db.AddMaterial(r.Context(), newMaterial.Material)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
