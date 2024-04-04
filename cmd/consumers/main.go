package main

import (
	"context"

	"github.com/NiltonMorais/gotaxi/internal/application/controller"
	"github.com/NiltonMorais/gotaxi/internal/application/usecase"
	gateway "github.com/NiltonMorais/gotaxi/internal/infra/gateway/memory"
	"github.com/NiltonMorais/gotaxi/internal/infra/queue"
	repository "github.com/NiltonMorais/gotaxi/internal/infra/repository/memory"
)

func main() {
	ctx := context.Background()
	queue := queue.NewRabbitMQAdapter("amqp://guest:guest@localhost:5672/")
	err := queue.Connect(ctx)
	if err != nil {
		panic(err)
	}
	accountRepository := repository.NewAccountMemoryRepository()
	emailGateway := gateway.NewEmailMemoryGateway()
	signupUseCase := usecase.NewSignupUseCase(accountRepository, emailGateway)
	queueController := controller.NewQueueController(queue, signupUseCase)
	err = queueController.ConsumerSignup(ctx)
	if err != nil {
		panic(err)
	}
	defer queue.Disconnect(ctx)
}
