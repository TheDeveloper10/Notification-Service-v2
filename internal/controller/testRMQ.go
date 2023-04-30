package controller

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type TestRMQ struct {
}

func (ctrl *TestRMQ) Get(request *amqp.Delivery) (any, bool) {
	return nil, true
}
