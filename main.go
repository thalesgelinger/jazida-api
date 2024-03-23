package main

import (
	"jazida-api/db"
	"jazida-api/midw"
	"jazida-api/services"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	db.CreateLoadsTable()

	mux.HandleFunc("GET /api/loads", midw.WithAdminAuth(services.GetLoads))
	mux.HandleFunc("POST /api/load", midw.WithLoaderAuth(services.SaveLoad))

	log.Println("Server Started")
	http.ListenAndServe("localhost:8080", mux)
}
