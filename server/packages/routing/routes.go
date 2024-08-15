package routing

import (
	"encoding/json"
	"net/http"

	"github.com/humanbojack/again/server/packages/db"
)

func CreateRoutes(router *http.ServeMux, database db.Database) {
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := database.GetAllTasks()
		if err != nil {
			http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})
}
