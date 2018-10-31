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

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Yet Another Todo App!")
}
