package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rmadamanchi/todo/internal/tasks"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")

	tasksApiRouter := router.PathPrefix("/api/tasks").Subrouter()
	tasks.RegisterApiHandlers(tasksApiRouter)

	tasksUiRouter := router.PathPrefix("/ui/tasks").Subrouter()
	tasks.RegisterUiHandlers(tasksUiRouter)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-gracefulStop
		fmt.Println("Shutting Down..")
		os.Exit(1)
	}()

	fmt.Println("Starting Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(writer, "Yet Another Todo App!")
}
