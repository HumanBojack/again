package routing

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/humanbojack/again/server/packages/db"
	"github.com/humanbojack/again/server/packages/templates"
)

type HtmlHandler struct {
	DB db.Database
}

func NewHtmlHandler(db db.Database) *HtmlHandler {
	return &HtmlHandler{DB: db}
}

func (h *HtmlHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	ti, err := db.TaskInputFromForm(r.PostForm)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse frequency: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	err = h.DB.CreateTask(&db.Task{TaskInput: ti})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to insert in db: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	tmpl := `
		<h1>task correctly inserted</h1>
	`
	t := template.Must(template.New("a").Parse(tmpl))
	t.Execute(w, nil)
}

func (h *HtmlHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Task"))
}

func (h *HtmlHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Task"))
}

func (h *HtmlHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Task"))
}

func (h *HtmlHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.DB.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFS(templates.FS, "html/tasks.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, tasks)
}
