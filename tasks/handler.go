package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var repository = NewRepository(MemoryMap)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("", handleGetTasks).Methods("GET")
	router.HandleFunc("", handlePostTask).Methods("POST")
	router.HandleFunc("/{id}", handleGetTask).Methods("GET")
}

func handleGetTasks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repository.all())
}

func handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 16)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	task := repository.get(int16(id))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func handlePostTask(w http.ResponseWriter, r *http.Request) {

}