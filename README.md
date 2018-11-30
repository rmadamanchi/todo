# todo
A simple todo application to learn Golang and more 

# Lab

### Create app folders

```bash
mkdir ~/src/github.com/<username>/todo
cd  ~/src/github.com/<username>/todo
```

### Hello World
Create `main.go`

```go
import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

Run main.go
```bash
go run main.go
```

Fix package
```go
package todo 
```

Fix package
```go
package main
```

Fix package
```go
go run main.go
```

### Dep Project with Mux

Install Dep
```bash
// can also install via curl or brew
go get -u github.com/golang/dep/cmd/dep
```

```bash
cd ~/src/github.com/<username>/todo
dep init
dep ensure https://github.com/gorilla/mux
```

```go
package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", handleHome).Methods("GET")

  fmt.Println("Starting Server on port 8000")
  http.ListenAndServe(":8000", router)
}

func handleHome(writer http.ResponseWriter, _ *http.Request) {
  writer.WriteHeader(http.StatusOK)
  fmt.Fprintf(writer, "Yet Another Todo App!")
}
```

### Model and GET Actions

Create a new package `tasks`

Create `tasks/model.go` to hold model structs

```go
package main

type Task struct {
	Id    int16  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
```

Add a GET handler for `/tasks` in `main.go`

```go
router.HandleFunc("/tasks", handleGetTasks).Methods("GET")

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
  writer.WriteHeader(http.StatusOK)
  json.NewEncoder(writer).Encode([]Task {
    Task {Id: 1, Title: "Get Milk"},
    Task {Id: 2, Title: "Get Bread"},
  })
}
```

Now run using

```bash
go run *.go
```

```bash
go build 
```

```bash
go install
```

Create a `tasks` package and move `model.go` into it

Create `tasks/handlers.go` and move handler logic into it

```go
package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rmadamanchi/todo/tasks"
	"net/http"
)

func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("", handleGetTasks).Methods("GET")
}

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode([]tasks.Task {
		tasks.Task {Id: 1, Title: "Get Milk"},
		tasks.Task {Id: 2, Title: "Get Bread"},
	})
}
```

Update `main.go` to use call `tasks.RegisterHandlers`

```go
tasksRouter := router.PathPrefix("/tasks").Subrouter()
tasks.RegisterHandlers(tasksRouter)
```

Extract an in-memory db with an array

```go
package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var db = [2]Task{
	Task{Id: 1, Title: "Get Bread"},
	Task{Id: 2, Title: "Get Milk"},
}


func RegisterHandlers(router *mux.Router) {
	router.HandleFunc("", handleGetTasks).Methods("GET")
}

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(db)
}
```

You can also initialize the array like this using the package init function
```go
var db [2]Task

func init() {
	db[0] = Task{Id: 1, Title: "Get Bread"}
	db[1] = Task{Id: 2, Title: "Get Bread"}
}
```


## POST Action
You can't add elements to array (you'd have to recreate and copy over). Use a slice instead.

Omitting the array length makes it a slice

```go
var db = []Task{
	Task{Id: 1, Title: "Get Bread"},
	Task{Id: 2, Title: "Get Milk"},
}
```

or using a package init method

```go
var db []Task

func init() {
	db = append(db, Task{Id: 1, Title: "Get Bread"})
	db = append(db, Task{Id: 2, Title: "Get Milk"})
}
```

Add a POST Action

```go
router.HandleFunc("", handlePostTask).Methods("POST")

func handlePostTask(writer http.ResponseWriter, request *http.Request) {
	var task Task
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&task); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(map[string]string{"error": "Invalid Request Body - " + err.Error())})
		writer.Write(response)
		return
	}

	db = append(db, task)
	writer.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(task)
	writer.Write(response)
}

```

Let's extract methods to read input body and send Responses

```go
func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	sendJson(writer, repository.all())
}

func handlePostTask(writer http.ResponseWriter, request *http.Request) {
	task, err := readBody(request)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "Invalid Request Body - "+err.Error())
		return
	}

	db = append(db, *task)
	sendJson(writer, task)
}

// return Task pointer since we need to return nil in case of error
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
```

Add a package level taskCounter to assign ids

```go
var taskCounter int16 = 1
```

Assign the id before saving a task in `handlePostTask`

```go
task.Id = taskCounter
taskCounter += 1

...
db = append(db, *task)
```

## Extract a repository

Create `tasks/repository.go`

```go
package tasks

type Repository interface {
	all() []Task
	get(id int16) Task
	update(t *Task)
	delete(id int16)
	create(t *Task)
}
```

Create memoryRepository in `tasks/repository.go` as a `Repository` implementation. Notice the pointer receiver

```go
type memoryMapRepository struct {
	db        map[int16]Task
	idCounter int16
}

func (repository memoryMapRepository) all() []Task {
	var tasks []Task
	for _, task := range repository.db {
		tasks = append(tasks, task)
	}
	return tasks
}

func (repository memoryMapRepository) get(id int16) Task {
	return repository.db[id]
}

func (repository memoryMapRepository) update(ask *Task) {
	repository.db[ask.Id] = *ask
}

func (repository memoryMapRepository) delete(id int16) {
	delete(repository.db, id)
}

func (repository *memoryMapRepository) create(task *Task) {
	task.Id = repository.idCounter
	repository.db[task.Id] = *task
	repository.idCounter += 1
}

```

Create a factory method to create a `Repository` instance 
```go
type RepositoryType int

const (
	MemoryMap RepositoryType = iota
)

func NewRepository(repositoryType RepositoryType) Repository {
	switch repositoryType {
	case MemoryMap:
		return &memoryMapRepository{db: make(map[int16]Task), idCounter: 1}
	default:
		return nil
	}
}
```

Create a test `tests/repository_test.go`

```go
package tasks

import "testing"

func TestMemoryMapRepository(t *testing.T) {
	repository := NewRepository(MemoryMap)

	repository.create(&Task{Title: "Get Milk"})
	repository.create(&Task{Title: "Get Bread"})
	repository.create(&Task{Title: "Fill Gas", Done: true})

	all := repository.all()
	assert(t, len(all), 3)

	task := repository.get(1)
	assert(t, task.Title, "Get Milk")
}

func assert(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
```

Run the tests using

```bash
go test
```


Update `handlers.go` to use the new repository

```go
var repository = NewRepository(MemoryMap)

func handleGetTasks(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	sendJson(w, repository.all())
}

func handlePostTask(writer http.ResponseWriter, request *http.Request) {
	task, err := readBody(request)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "Invalid Request Body - "+err.Error())
		return
	}

	repository.create(task)
	sendJson(writer, task)
}
```