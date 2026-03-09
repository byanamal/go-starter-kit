package main

import (
	"context"
	"log"

	"base-api/internal/pkg/config"
	"base-api/internal/pkg/db"
	"base-api/internal/pkg/seeder"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbConfig := config.LoadDatabase()
	dbConn := db.Init(dbConfig.DSN())

	ctx := context.Background()

	if err := seeder.Run(ctx, dbConn); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database seeded successfully")
}
