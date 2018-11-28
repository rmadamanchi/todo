# todo
A simple todo application to learn Golang and more 

# Lab

### Create app folders

```
mkdir ~/src/github.com/<username>/todo
cd  ~/src/github.com/<username>/todo
```

### Hello World
`vi main.go`

```
import "fmt"

func main() {
  fmt.Println("Hello World")
}
```

Run main.go
```
go run main.go
```

Fix package
```
package todo 
```

Fix package
```
package main
```

Fix package
```
go run main.go
```

### Dep Project with Mux

Install Dep
```
// can also install via curl or brew
go get -u github.com/golang/dep/cmd/dep
```

```
cd ~/src/github.com/<username>/todo
dep init
dep ensure https://github.com/gorilla/mux
```

```
package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", homeHandler).Methods("GET")

  fmt.Println("Starting Server on port 8000")
  http.ListenAndServe(":8000", router)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Yet Another Todo App!")
}
```

### Model and REST Actions