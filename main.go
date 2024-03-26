package main

import (
	"fmt"
	"jazida-api/db"
	"jazida-api/midw"
	"jazida-api/services"
	"jazida-api/views"
	"log"
	"net/http"
	"os"
)

func main() {
	db.CreateLoadsTable()
	db.CreateClientsTable()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/loads", midw.WithAdminAuth(services.GetLoads))
	mux.HandleFunc("POST /api/load", midw.WithLoaderAuth(services.SaveLoad))
	mux.HandleFunc("POST /api/signature", midw.Cors(midw.WithLoaderAuth(services.SaveSignature)))

	mux.HandleFunc("/", views.Home)
	mux.HandleFunc("/clients", views.Config)
	mux.HandleFunc("GET /new-client", views.NewFormClient)
	mux.HandleFunc("POST /new-client", views.AddClient)

	log.Println("Server Started")

	cwd, _ := os.Getwd()
	signaturesPath := fmt.Sprintf("%s/signatures", cwd)
	fs := http.FileServer(http.Dir(signaturesPath))

	mux.Handle("/api/signatures/", http.StripPrefix("/api/signatures/", fs))
	http.ListenAndServe("localhost:8080", mux)
}
