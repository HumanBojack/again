package main

import (
	"log"
	"net/http"
	"os"

	"github.com/humanbojack/again/server/packages/middlewares"
)

func main() {
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
