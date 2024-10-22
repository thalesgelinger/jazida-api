package server

import (
	"encoding/json"
	"fmt"
	"jazida-api/internal/handler"
	"jazida-api/internal/infra/db"
	"jazida-api/internal/middleware"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
	"golang.org/x/net/websocket"
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

	cwd, _ := os.Getwd()
	signaturesPath := fmt.Sprintf("%s/signatures", cwd)
	fs := http.FileServer(http.Dir(signaturesPath))

	s.router.Handle("/api/signatures/", http.StripPrefix("/api/signatures/", fs))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	handler := c.Handler(s.router)

	return http.ListenAndServe(s.listenAddr, handler)
}

func (s *Server) setupLoadRoutes() {

	socket := handler.NewSocket()
	lh := handler.NewLoadHandler(queries, socket)

	s.router.Handle("/new-load-added", websocket.Handler(func(w *websocket.Conn) {
		socket.AddClient(w)
		defer socket.RemoveClient(w)
		for {
			time.Sleep(1 * time.Second)
		}
	}))
	s.router.HandleFunc("GET /api/loads", midw.WithAdminAuth(lh.GetLoads))
	s.router.HandleFunc("POST /api/load", midw.WithLoaderAuth(lh.SaveLoad))
	s.router.HandleFunc("POST /api/signature", midw.Cors(midw.WithLoaderAuth(lh.SaveSignature)))
}

func (s *Server) setupClientsRoutes() {

	ch := handler.NewClientHandler(queries)

	s.router.HandleFunc("GET /api/clients", midw.WithLoaderAuth(ch.GetClients))
	s.router.HandleFunc("POST /api/clients", midw.WithAdminAuth(ch.CreateClient))
	s.router.HandleFunc("POST /api/clients/{id}/plates", midw.WithAdminAuth(ch.CreatePlate))
	s.router.HandleFunc("GET /api/materials", midw.WithLoaderAuth(ch.GetMaterials))
	s.router.HandleFunc("POST /api/materials", midw.WithAdminAuth(ch.CreateMaterial))
}
