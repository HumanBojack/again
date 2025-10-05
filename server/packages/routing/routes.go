package routing

import (
	"net/http"
)

func CreateRoutes(router *http.ServeMux, h Handler) {
	router.Handle("GET /", http.HandlerFunc(h.GetTasks))
	router.Handle("POST /task", http.HandlerFunc(h.CreateTask))
	router.Handle("GET /task/{id}", http.HandlerFunc(h.GetTask))
	router.Handle("POST /task/{id}", http.HandlerFunc(h.UpdateTask))
	router.Handle("DELETE /task/{id}", http.HandlerFunc(h.DeleteTask))
}
