# cli sample

## description

This is a sample project to demonstrate how to use the cli module.

## init

```bash
go get -u github.com/spf13/cobra
go mod tidy
```

## run

```bash
go run cli.go
```

## build

```bash
go build -o cli cli.go
```

use `go clean` to remove the binary

## install

install the binary to the $GOPATH/bin directory

```bash
go install
```

uninstall the binary from the $GOPATH/bin directory, when want to remove the binary

