package services

import (
	"encoding/json"
	"jazida-api/db"
	"jazida-api/entities"
	"log"
	"net/http"
)

func GetLoads(w http.ResponseWriter, r *http.Request) {
	loads, err := db.GetLoads()
	if err != nil {
		log.Fatal("Error getting loads")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(loads)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SaveLoad(w http.ResponseWriter, r *http.Request) {

	var newLoad entities.Load
	if err := json.NewDecoder(r.Body).Decode(&newLoad); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	db.SaveLoad(&newLoad)
}
