package main

import (
	"notification-service/internal/api"
	"notification-service/internal/api/middleware"
	"notification-service/internal/client"
	"notification-service/internal/config"
	"notification-service/internal/util"
)

func main() {
	// Initializations
	config.InitConfigs()
	client.InitClients()
	middleware.InitMiddlewares()

	// Handlers
	var (
		hasHTTPHandler = false

		fiberApp = api.NewFiberApp()
	)

	// HTTP REST V1
	if config.Master.Service.APIs.Has(config.HTTP_REST_V1_API) {
		hasHTTPHandler = true

		api.SetUpRESTV1(fiberApp)
	}

	// RabbitMQ V1
	if config.Master.Service.APIs.Has(config.RABBITMQ_V1_API) {
		api.SetUpRabbitMQV1()
	}

	// Starting HTTP server
	if hasHTTPHandler {
		util.Logger.Info().Msg("HTTP Server is ON")
		err := fiberApp.Listen(config.Master.HTTP.Address)
		if err != nil {
			util.Logger.Error().Msg(err.Error())
			util.Logger.Panic().Msg("HTTP Server failed!")
		}
	}
}
