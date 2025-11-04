build:
	@go build -o bin/go-cli-game cmd/main.go

run: build
	@./bin/go-cli-game

dev:
	@go run cmd/main.go