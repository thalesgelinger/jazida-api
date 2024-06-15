package main

import (
	"context"
	"flag"
	"fmt"
	"jazida-api/internal/server"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	port := flag.String("port", ":8080", "The server running port")

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	ctx := context.Background()

	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to parse database URL: %s", err.Error()))
		return
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		slog.Error(fmt.Sprintf("DB not connected: %s", err.Error()))
		return
	}
	defer pool.Close()
	slog.Info("DB connected")

	server := server.NewServer(*port, pool)

	slog.Info(fmt.Sprintf("Server running at port: %s", *port))

	log.Fatal(server.Start(), pool)
}
