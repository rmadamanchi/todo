package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rmadamanchi/todo/internal/tasks"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")

	tasksApiRouter := router.PathPrefix("/api/tasks").Subrouter()
	tasks.RegisterApiHandlers(tasksApiRouter)

	tasksUiRouter := router.PathPrefix("/ui/tasks").Subrouter()
	tasks.RegisterUiHandlers(tasksUiRouter)

	fmt.Println("Starting Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func homeHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(writer, "Yet Another Todo App!")
}
