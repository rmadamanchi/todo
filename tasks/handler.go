package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var repository, _ = NewRepository(MemoryMap)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("", handleGetTasks).Methods("GET")
	router.HandleFunc("", handlePostTask).Methods("POST")
	router.HandleFunc("/{id}", handleGetTask).Methods("GET")
}

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	sendJson(writer, repository.all())
}

func handleGetTask(writer http.ResponseWriter, _ *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 16)
	if err == nil {
		sendError(writer, http.StatusBadRequest, "Bad Id - " + vars["id"])
	}
	task := repository.get(int16(id))

	sendJson(w, task)
	json.NewEncoder(w).Encode(task)
}

func handlePostTask(writer http.ResponseWriter, request *http.Request) {
	task, err := readBody(request)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "Invalid Request Body - " + err.Error())
		return
	}

	repository.create(task)
	sendJson(writer, task)
}

func readBody(request *http.Request) (*Task, error) {
	var task Task
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func sendJson(writer http.ResponseWriter, body interface{}) {
	sendResponse(writer, http.StatusOK, body)
}

func sendError(writer http.ResponseWriter, code int, message string) {
	sendResponse(writer, code, map[string]string{"error": message})
}

func sendResponse(writer http.ResponseWriter, code int, body interface{}) {
	response, _ := json.Marshal(body)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(response)
}
