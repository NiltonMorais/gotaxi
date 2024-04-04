package main

import (
	"context"
	"log"

	"github.com/NiltonMorais/gotaxi/internal/infra/queue"
)

func main() {
	ctx := context.Background()
	rabbitmq := queue.NewRabbitMQAdapter("amqp://guest:guest@localhost:5672/")
	err := rabbitmq.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer rabbitmq.Disconnect(ctx)

	rabbitmq.Publish(ctx, "test", "Hello World!")
}
