### Create app folders

```bash
mkdir ~/go/src/github.com/<username>/todo
cd  ~/go/src/github.com/<username>/todo
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