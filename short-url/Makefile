PROGRAM = "./cmd/url.go"

all:
	make swag
	make fmt
	make run

swag:
	cp ./cmd/url.go ./main.go
	swag fmt
	swag init
	mv ./main.go ./cmd/url.go

init:
	go mod init short-url
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/echo-swagger
	docker run -d -p 5672:5672 --hostname my-rabbit --name some-rabbit rabbitmq:3

test:
	go clean -testcache
	go test -v ./... -cover

fmt:
	go mod tidy
	go fmt ./...

build:
	make fmt
	go build -o url $(PROGRAM)

run:
	go run $(PROGRAM)