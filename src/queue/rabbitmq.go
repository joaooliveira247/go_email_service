package queue

import (
	"fmt"
	"log/slog"

	"github.com/joaooliveira247/go_email_service/src/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConnection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (rc *RabbitConnection) declareQueues() error {

	// DLQ declaration
	_, err := rc.Channel.QueueDeclare(
		config.EmailDLQQueue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed when try declare DLQ: %w", err)
	}

	// Retry Queue
	retryArgs := amqp.Table{
		"x-dead-letter-exchange":    config.RabbitExchange,
		"x-dead-letter-routing-key": config.EmailQueue,
		"x-message-ttl":             config.EmailRetryTTL,
	}

	_, err = rc.Channel.QueueDeclare(
		config.EmailRetryQueue,
		true,
		false,
		false,
		false,
		retryArgs,
	)

	if err != nil {
		return fmt.Errorf("Failed when try declare RetryQueue: %w", err)
	}

	// Main Queue
	_, err = rc.Channel.QueueDeclare(
		config.EmailQueue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed when try declare MainQueue: %w", err)
	}

	return nil
}

func (rc *RabbitConnection) Close() {
	if rc.Channel != nil {
		rc.Channel.Close()
	}
	if rc.Conn != nil {
		rc.Conn.Close()
	}
	slog.Info("Conections with RabbitMQ closed")
}
