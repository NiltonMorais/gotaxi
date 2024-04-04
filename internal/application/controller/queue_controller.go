package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/NiltonMorais/gotaxi/internal/application/usecase"
	"github.com/NiltonMorais/gotaxi/internal/infra/queue"
)

type QueueController struct {
	queue         queue.Queue
	signupUseCase *usecase.SignupUseCase
}

func NewQueueController(queue queue.Queue, signupUseCase *usecase.SignupUseCase) *QueueController {
	return &QueueController{
		queue:         queue,
		signupUseCase: signupUseCase,
	}
}

func (qc *QueueController) ConsumerSignup(ctx context.Context) error {
	callback := func(message queue.QueueMessage) error {
		log.Printf("Message received: %s", message.Body)
		var data SignupData
		err := json.Unmarshal(message.Body, &data)
		if err != nil {
			return errors.New("Erro ao decodificar JSON:" + err.Error())
		}
		err = qc.signupUseCase.Execute(ctx, data.Name, data.Email, data.Document, data.CarPlate, data.IsPassenger, data.IsDriver)
		if err != nil {
			return errors.New("Erro ao executar o caso de uso:" + err.Error())
		}
		log.Printf("Processou a message: %s", message.Body)
		return nil
	}
	return qc.queue.Consume(ctx, "signup", callback)
}

type SignupData struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Document    string `json:"document"`
	CarPlate    string `json:"car_plate"`
	IsPassenger bool   `json:"is_passenger"`
	IsDriver    bool   `json:"is_driver"`
}
