package routing

import (
	"html/template"
	"net/http"

	"github.com/humanbojack/again/server/packages/db"
)

type HtmlHandler struct {
	DB db.Database
}

func NewHtmlHandler(db db.Database) *HtmlHandler {
	return &HtmlHandler{DB: db}
}

func (h *HtmlHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Task"))
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
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Tasks</title>
	</head>
	<body>
		<h1>Tasks</h1>
		<ul>
		{{range .}}
			<li>{{.Title}}</li>
		{{end}}
		</ul>
	</body>
	</html>
	`
	t := template.Must(template.New("tasks").Parse(tmpl))
	t.Execute(w, tasks)
}
