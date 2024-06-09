run: build
	@./bin/my-community

run-dev:
	@go run cmd/api/main.go

build:
	@go build -o bin/my-community cmd/api/main.go

