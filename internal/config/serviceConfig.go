package config

import "notification-service/internal/util"

type MasterConfig struct {
	Service struct {
		APIs              util.Strings `yaml:"apis"`
		NotificationTypes util.Strings `yaml:"notification_types"`
	} `yaml:"service"`
}
