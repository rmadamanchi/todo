package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHanders(router *mux.Router) {
	router.HandleFunc("", HandleGetTasks).Methods("GET")
}

func HandleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode([]Task{
		{Id: 1, Title: "Get Milk", Done: false},
		{Id: 1, Title: "Get Milk", Done: false},
	})
}
