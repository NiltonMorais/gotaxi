package queue

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	uri  string
	conn *amqp.Connection
}

func NewRabbitMQAdapter(uri string) *RabbitMQAdapter {
	return &RabbitMQAdapter{
		uri: uri,
	}
}

func (r *RabbitMQAdapter) Connect(ctx context.Context) error {
	conn, err := amqp.Dial(r.uri)
	if err != nil {
		return err
	}
	r.conn = conn
	return nil
}

func (r *RabbitMQAdapter) Disconnect(ctx context.Context) error {
	return r.conn.Close()
}

func (r *RabbitMQAdapter) Publish(ctx context.Context, queueName, message string) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s\n", message)
	return nil
}

func (r *RabbitMQAdapter) Consume(ctx context.Context, queueName string, callback func(QueueMessage) error) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err := callback(QueueMessage{Body: d.Body})
			if err != nil {
				log.Printf("Error processing message: %s", err)
			} else {
				d.Ack(false)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	var forever chan struct{}
	<-forever
	return nil
}
