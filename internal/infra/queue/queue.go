package queue

import "context"

type QueueMessage struct {
	Body []byte
}

type Queue interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Consume(ctx context.Context, queueName string, callback func(QueueMessage) error) error
	Publish(ctx context.Context, queueName, message string) error
}
