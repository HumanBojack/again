package routing

import (
	"encoding/json"
	"net/http"

	"github.com/humanbojack/again/server/packages/db"
)

type JSONHandler struct {
	DB db.Database
}

func NewJSONHandler(db db.Database) *JSONHandler {
	return &JSONHandler{DB: db}
}

func (h *JSONHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task db.TaskInput
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	newTask := db.Task{
		TaskInput: task,
	}
	err = h.DB.CreateTask(&newTask)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (h *JSONHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := h.DB.GetTask(id)
	if err != nil {
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *JSONHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var task db.TaskInput
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	newTask := db.Task{
		TaskInput: task,
	}

	err = h.DB.UpdateTask(id, &newTask)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func (h *JSONHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.DB.DeleteTask(id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *JSONHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.DB.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
