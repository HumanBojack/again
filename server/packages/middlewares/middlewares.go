package middlewares

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.Handler, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f.ServeHTTP)
	}
	return f.ServeHTTP
}

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("Request: %s %s %s", r.Method, r.URL.Path, time.Since(start))
	}
}

func ApiKeyMiddlewareGenerator(apiKey string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if apiKey != "" && r.Header.Get("X-API-KEY") != apiKey {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next(w, r)
		}
	}
}
