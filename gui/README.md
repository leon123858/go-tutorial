# GUI Sample

this is a quick sample of a GUI with Golang

## How to run

use `go run gui.go` to run the program

## create this project

```bash
go install gioui.org/cmd/gogio@latest
go mod init gui
go get gioui.org/example/kitchen
gogio -target js gioui.org/example/kitchen
```

write below code in `gui.go`

```go
package main

import "net/http"

func main() {
	// print on localhost:8080
	println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("kitchen")))
	if err != nil {
		return
	}
}
```

## How to clean

```bash
rm -rf go.mod go.sum kitchen
```