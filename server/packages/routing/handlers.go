package routing

import (
	"encoding/json"
	"net/http"

	"github.com/humanbojack/again/server/packages/db"
)

type Handler struct {
	DB db.Database
}

func NewHandler(db db.Database) *Handler {
	return &Handler{DB: db}
}

type TasksHandler interface {
	GetTasks(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.DB.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
