package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"calendar_app/internal/api"
	"calendar_app/internal/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbAddr := os.Getenv("DB_CONN")

	conn, err := pgxpool.New(ctx, dbAddr)
	if err != nil {
		log.Fatalf("unable to connect to db: %s", err)
	}

	dbAdapter := postgres.New(conn)

	server := api.New(dbAdapter)

	addr := os.Getenv("APP_ADDR")
	if err := server.Run(addr); err != nil {
		log.Fatalf("server.Run: %s", err.Error())
	}
}
