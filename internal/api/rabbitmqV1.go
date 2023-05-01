package api

import (
	"notification-service/internal/client"
	"notification-service/internal/util"
)

func SetUpRabbitMQV1() {
	client.RabbitMQ.Consume("test", testRMQCtrl.Get)

	client.RabbitMQ.Consume("write_template", templateRMQCtrl.Write)
	client.RabbitMQ.Consume("get_template", templateRMQCtrl.Get)
	client.RabbitMQ.Consume("delete_template", templateRMQCtrl.Delete)

	client.RabbitMQ.Consume("send_notifications", notificationRMQCtrl.Send)

	util.Logger.Info().Msg("Initialized RabbitMQ V1 routes")
}
