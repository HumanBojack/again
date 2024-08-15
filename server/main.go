package main

import (
	"log"
	"net/http"
	"os"

	"github.com/humanbojack/again/server/packages/db"
	"github.com/humanbojack/again/server/packages/middlewares"
	"github.com/humanbojack/again/server/packages/routing"
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
	jr := http.NewServeMux()
	hr := http.NewServeMux()

	app_db := db.NewGormDB(database)
	jh := routing.NewJsonHandler(app_db)
	hh := routing.NewHtmlHandler(app_db)
	routing.CreateRoutes(jr, jh)
	routing.CreateRoutes(hr, hh)

	// Middlewares
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("API_KEY is not set, the server will not be protected by an API key")
	}
	apiKeyMiddleware := middlewares.ApiKeyMiddlewareGenerator(apiKey)

	contentTypeMiddleware := middlewares.ContentTypeMiddlewareGenerator(
		map[string]http.Handler{
			"application/json": jr,
			"text/html":        hr,
		},
	)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewares.Chain(jr, contentTypeMiddleware, apiKeyMiddleware, middlewares.LogMiddleware),
	}

	server.ListenAndServe()
}
