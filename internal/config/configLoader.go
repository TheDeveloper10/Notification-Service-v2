package config

import (
	"encoding/json"
	"io/ioutil"
	"notification-service/internal/util"
	"os"
)

func loadConfig(fileName string, out any) bool {
	file, err := os.Open(fileName)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to open config file %s", fileName)
		return false
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to read config file %s", fileName)
		return false
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msgf("Failed to unmarshal config file %s", fileName)
		return false
	}

	util.Logger.Info().Msgf("Loaded config from %s", fileName)
	return true
}
