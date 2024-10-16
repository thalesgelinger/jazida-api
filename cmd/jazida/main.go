package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"jazida-api/internal/server"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	godotenv.Load()

	port := flag.String("port", ":8080", "The server running port")

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	ctx := context.Background()

	// Open SQLite database connection
	db, err := sql.Open("sqlite3", "./db/db.sqlite")
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to open SQLite database: %s", err.Error()))
		return
	}
	defer db.Close()

	// Test the connection
	err = db.PingContext(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("DB not connected: %s", err.Error()))
		return
	}
	slog.Info("DB connected")

	// Pass the SQLite db connection to the server
	server := server.NewServer(*port, db)

	slog.Info(fmt.Sprintf("Server running at port: %s", *port))

	log.Fatal(server.Start(), db)
}

