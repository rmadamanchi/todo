package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("", handleGetTasks).Methods("GET")
}

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode([]Task{
		{Id: 1, Title: "Get Milk", Done: false},
		{Id: 1, Title: "Get Milk", Done: false},
	})
}
