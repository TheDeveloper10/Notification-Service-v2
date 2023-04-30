package api

import (
	"notification-service/internal/client"
	"notification-service/internal/controller"
)

func SetUpRabbitMQV1() {
	var (
		testCtrl = controller.NewTestRMQController()
	)

	client.RabbitMQ.Consume("test", testCtrl.Get)
}
