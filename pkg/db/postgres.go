package db

import (
	"context"
	"log"

	"bingo-backend/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitPostgres() {

	var err error

	Pool, err = pgxpool.New(context.Background(), config.App.DatabaseURL)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	err = Pool.Ping(context.Background())

	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Connected to Neon PostgreSQL")
}