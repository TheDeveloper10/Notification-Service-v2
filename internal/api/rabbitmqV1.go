package api

import (
	"notification-service/internal/client"
	"notification-service/internal/util"
)

func SetUpRabbitMQV1() {
	client.RabbitMQ.Consume("test", testRMQCtrl.Get)

	util.Logger.Info().Msg("Initialized RabbitMQ V1 routes")
}
