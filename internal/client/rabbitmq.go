package client

import (
	"context"
	"encoding/json"
	"fmt"
	"notification-service/internal/config"
	"notification-service/internal/util"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQClient(conf *config.RabbitMQConfig) *RabbitMQClient {
	url := fmt.Sprintf("amqp://%s:%s@%s", conf.Username, conf.Password, conf.Host)

	connection, err := amqp.Dial(url)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("Failed to initialize RabbitMQ client")
		return nil
	}

	channel, err := connection.Channel()
	if err != nil {
		connection.Close()
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("Failed to initialize RabbitMQ channel")
		return nil
	}

	c := RabbitMQClient{
		connection: connection,
		channel:    channel,
	}

	util.Logger.Info().Msg("Initialized RabbitMQ Client")
	return &c
}

type RabbitMQClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

type RabbitMQRequests <-chan amqp.Delivery
type RabbitMQHandler func(*amqp.Delivery) (any, bool)

func (rmq *RabbitMQClient) Close() {
	rmq.channel.Close()
	rmq.connection.Close()
}

func (rmq *RabbitMQClient) Consume(queueName string, handler RabbitMQHandler) {
	_, err := rmq.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to initialize RabbitMQ queue %s", queueName)
		return
	}

	requests, err := rmq.channel.Consume(
		queueName,
		"", false, false, false, false, nil,
	)
	if err != nil {
		rmq.Close()
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("Failed to initialize RabbitMQ channel")
		return
	}

	go rmq.handleRequests(requests, handler)
}

func (rmq *RabbitMQClient) handleRequests(requests RabbitMQRequests, handler RabbitMQHandler) {
	for request := range requests {
		go rmq.processRequest(request, handler)
	}
}

func (rmq *RabbitMQClient) processRequest(request amqp.Delivery, handler RabbitMQHandler) {
	resp, ack := handler(&request)

	if ack {
		err := request.Ack(false)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			return
		}

		if resp != nil && request.ReplyTo != "" {
			rmq.publishMessage(resp, request.ReplyTo, request.CorrelationId)
		}
	}
}

func (rmq *RabbitMQClient) publishMessage(message any, queueName string, correlationId string) {
	contentType := ""
	var body []byte

	if _, ok := message.(string); ok {
		body = []byte(message.(string))
		contentType = "application/text"
	} else {
		temp, err := json.Marshal(message)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			return
		} else {
			contentType = "application/json"
			body = temp
		}
	}

	err := rmq.channel.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode:  amqp.Persistent,
			Timestamp:     time.Now(),
			ContentType:   contentType,
			CorrelationId: correlationId,
			Body:          body,
		},
	)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
	}
}
