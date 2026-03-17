BINARY_NAME=app
MAIN_PATH=./app/main.go

.PHONY: all build run clean test lint

all: build

build:
	@echo "Compilazione in corso..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	@go run $(MAIN_PATH)

clean:
	@echo "Pulizia..."
	@rm -f $(BINARY_NAME)
	@go clean

test:
	@go test -v ./...

lint:
	@golangci-lint run
