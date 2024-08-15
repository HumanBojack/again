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
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetTasks(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.TaskInput
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = h.DB.CreateTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
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
