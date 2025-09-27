run:
	go run cmd/go-api/main.go

build:
	go build -o bin/go-api cmd/go-api/main.go

run-build:
	./bin/go-api