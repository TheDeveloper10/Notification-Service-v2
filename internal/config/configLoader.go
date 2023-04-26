package config

import (
	"io"
	"notification-service/internal/util"
	"os"

	"gopkg.in/yaml.v2"
)

func loadYAML(fileName string, out any) bool {
	file, err := os.Open(fileName)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to load config from %s", fileName)
		return false
	}

	data, err := io.ReadAll(file)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to load config from %s", fileName)
		return false
	}

	err = yaml.Unmarshal(data, out)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to load config from %s", fileName)
		return false
	}

	util.Logger.Info().Msgf("Loaded config from %s", fileName)
	return true
}
