build: 
	@go build -o bin/gotaxi cmd/api/main.go

run: build
	@./bin/gotaxi

test:
	@go test -v ./...

producer:
	@go run internal/infra/queue/rabbitmq/producer/main.go

consumer:
	@go run cmd/consumers/main.go
	#@go run internal/infra/queue/rabbitmq/consumer/main.go

