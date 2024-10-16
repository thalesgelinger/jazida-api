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

	s.setupViews()

	cwd, _ := os.Getwd()
	signaturesPath := fmt.Sprintf("%s/signatures", cwd)
	fs := http.FileServer(http.Dir(signaturesPath))

	s.router.Handle("/api/signatures/", http.StripPrefix("/api/signatures/", fs))

	return http.ListenAndServe(s.listenAddr, s.router)
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
	s.router.HandleFunc("GET /api/loads", withCors(midw.WithAdminAuth(lh.GetLoads)))
	s.router.HandleFunc("POST /api/load", midw.WithLoaderAuth(lh.SaveLoad))
	s.router.HandleFunc("POST /api/signature", midw.Cors(midw.WithLoaderAuth(lh.SaveSignature)))
}

func (s *Server) setupClientsRoutes() {

	ch := handler.NewClientHandler(queries)

	s.router.HandleFunc("GET /api/clients", midw.WithLoaderAuth(ch.GetClients))
}

func (s *Server) setupViews() {

	fs := http.FileServer(http.Dir("./web/dist"))

	s.router.Handle("/", fs)
}

func withCors(handleFunc func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		handleFunc(w, r)
	}
}
