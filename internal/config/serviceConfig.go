package config

import "notification-service/internal/util"

type MasterConfig struct {
	Service struct {
		APIs              util.Strings `yaml:"apis"`
		NotificationTypes util.Strings `yaml:"notification_types"`
		AllowedLanguages  util.Strings `yaml:"allowed_languages"`
	} `yaml:"service"`

	HTTP struct {
		Address string `yaml:"address"`
	} `yaml:"http"`
}
