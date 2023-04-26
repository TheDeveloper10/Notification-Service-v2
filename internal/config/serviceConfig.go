package config

import "notification-service/internal/util"

type MasterConfig struct {
	Service  ServiceConfig  `yaml:"service"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database DatabaseConfig `yaml:"database"`
}

type ServiceConfig struct {
	APIs              util.Strings `yaml:"apis"`
	NotificationTypes util.Strings `yaml:"notification_types"`
	AllowedLanguages  util.Strings `yaml:"allowed_languages"`
}

type HTTPConfig struct {
	Address string `yaml:"address"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	PoolSize int    `yaml:"pool_size"`
}
