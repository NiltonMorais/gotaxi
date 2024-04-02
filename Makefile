build: 
	@go build -o bin/gotaxi cmd/api/main.go

run: build
	@./bin/gotaxi

test:
	@go test -v ./...
