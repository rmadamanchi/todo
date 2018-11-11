package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var repository = NewRepository(MemoryMap)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("", handleGetTasks).Methods("GET")
}

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(repository.all())
}
