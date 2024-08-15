package routing

import (
	"net/http"
)

func CreateRoutes(router *http.ServeMux, h TasksHandler) {
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.Handle("POST /task", http.HandlerFunc(h.CreateTask))
	router.Handle("GET /tasks", http.HandlerFunc(h.GetTasks))
}
