package main

import (
	"log"
	"net/http"
	"os"

	"github.com/humanbojack/again/server/packages/db"
	"github.com/humanbojack/again/server/packages/middlewares"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to database
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")

	// Migrate the schema
	database.AutoMigrate(&db.Task{})
	log.Println("Migrated the schema")

	// Web router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Middlewares
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("API_KEY is not set, the server will not be protected by an API key")
	}
	apiKeyMiddleware := middlewares.ApiKeyMiddlewareGenerator(apiKey)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewares.Chain(router, apiKeyMiddleware, middlewares.LogMiddleware),
	}

	server.ListenAndServe()
}
