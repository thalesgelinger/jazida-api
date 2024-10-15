package server

import (
	"context"
	"encoding/json"
	"fmt"
	"jazida-api/internal/handler"
	"jazida-api/internal/infra/db"
	"jazida-api/internal/middleware"
	"jazida-api/views"
	"net/http"
	"os"
	"time"
)

type Server struct {
	listenAddr string
	conn       db.DBTX
	router     *http.ServeMux
}

func NewServer(listenAddr string, conn db.DBTX) *Server {
	return &Server{
		listenAddr: listenAddr,
		conn:       conn,
	}
}

var queries *db.Queries

func (s *Server) Start() error {

	queries = db.New(s.conn)

	s.router = http.NewServeMux()

	s.router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		type Hello struct {
			Message string `json:"message"`
		}

		json.NewEncoder(w).Encode(Hello{
			Message: "jazida api is up and running",
		})
	})

	s.setupLoadRoutes()
	s.setupClientsRoutes()

	s.setupViews()

	cwd, _ := os.Getwd()
	signaturesPath := fmt.Sprintf("%s/signatures", cwd)
	fs := http.FileServer(http.Dir(signaturesPath))

	s.router.Handle("/api/signatures/", http.StripPrefix("/api/signatures/", fs))
	return http.ListenAndServe(s.listenAddr, s.router)
}

func (s *Server) setupLoadRoutes() {
	newLoadsChannel := make(chan string)

	lh := handler.NewLoadHandler(queries)

	s.router.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// Set headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create a context for handling client disconnection
		_, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send data to the client
		go func() {
			for data := range newLoadsChannel {
				fmt.Fprintf(w, "data: %s\n\n", data)
				w.(http.Flusher).Flush()
			}
		}()

		// Simulate sending data periodically
		for {
			newLoadsChannel <- time.Now().Format(time.TimeOnly)
			time.Sleep(1 * time.Second)
		}
	})
	s.router.HandleFunc("GET /api/loads", midw.WithAdminAuth(lh.GetLoads))
	s.router.HandleFunc("POST /api/load", midw.WithLoaderAuth(lh.SaveLoad))
	s.router.HandleFunc("POST /api/signature", midw.Cors(midw.WithLoaderAuth(lh.SaveSignature)))
}

func (s *Server) setupClientsRoutes() {

	ch := handler.NewClientHandler(queries)

	s.router.HandleFunc("GET /api/clients", midw.WithLoaderAuth(ch.GetClients))
}

func (s *Server) setupViews() {

	vh := views.NewViewHandler(queries)

	s.router.HandleFunc("/", vh.Home)
	s.router.HandleFunc("/clients", vh.Config)
	s.router.HandleFunc("GET /new-client", vh.NewFormClient)
	s.router.HandleFunc("POST /new-client", vh.AddClient)
}
