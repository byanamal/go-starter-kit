package main

import (
	"base-api/internal/pkg/config"
	"base-api/internal/pkg/db"
	"base-api/internal/pkg/logger"
	"base-api/internal/server"

	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	logger.Setup()

	dbConfig := config.LoadDatabase()
	dbConn := db.Init(dbConfig.DSN())


	mux := http.NewServeMux()

	srv := server.NewServer(mux, dbConn)

	log.Println("Server running on http://localhost:8080")
	if err := srv.HttpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
