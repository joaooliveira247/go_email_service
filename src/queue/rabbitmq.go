package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConnection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}


