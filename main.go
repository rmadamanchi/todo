package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rmadamanchi/todo/tasks"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")

	tasksRouter := router.PathPrefix("/tasks").Subrouter()
	tasks.RegisterHandlers(tasksRouter)

	fmt.Println("Starting Server on port 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func homeHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Yet Another Todo App!")
}
