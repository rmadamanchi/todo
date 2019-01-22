package tasks

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func RegisterUiHandlers(router *mux.Router) {
	router.HandleFunc("", handleTasksPage).Methods("GET")
}

func handleTasksPage(writer http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/tasks.html"))
	tasks := repository.all()
	_ = tmpl.ExecuteTemplate(writer, "layout", tasks)
}
