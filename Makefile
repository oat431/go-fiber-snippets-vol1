.PHONY: build run test clean tidy

APP_NAME = api
MAIN     = cmd/api/main.go
BIN_DIR  = bin

build:
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN)

run:
	@go run $(MAIN)

test:
	@go test -v -count=1 ./...

test-cover:
	@go test -v -count=1 -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

clean:
	@go clean
	@rm -rf $(BIN_DIR) coverage.out coverage.html

tidy:
	@go mod tidy

lint:
	@golangci-lint run ./...
