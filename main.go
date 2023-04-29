package main

import (
	"notification-service/internal/api"
	"notification-service/internal/client"
	"notification-service/internal/config"
	"notification-service/internal/util"
	"sync"
)

func main() {
	// Initializations
	config.InitConfigs()
	client.InitClients()

	// Handlers
	var (
		hasHTTPHandler = false
		fiberApp       = api.NewFiberApp()
		wg             = sync.WaitGroup{}
	)

	// HTTP REST V1
	if config.Master.Service.APIs.Has(config.HTTP_REST_V1_API) {
		hasHTTPHandler = true

		api.SetUpRESTV1(fiberApp)
	}

	// HTTP REST V2
	//

	// Starting handlers
	if hasHTTPHandler {
		wg.Add(1)
		go func() {
			defer wg.Done()

			util.Logger.Info().Msg("HTTP Server is ON")
			err := fiberApp.Listen(config.Master.HTTP.Address)
			if err != nil {
				util.Logger.Error().Msg(err.Error())
				util.Logger.Panic().Msg("HTTP Server failed!")
			}
		}()
	}

	wg.Wait()
}
