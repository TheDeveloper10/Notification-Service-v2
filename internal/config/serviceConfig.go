package config

import "notification-service/internal/util"

type MasterConfig struct {
	Service  ServiceConfig  `yaml:"service"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database DatabaseConfig `yaml:"database"`
	Mail     MailConfig     `yaml:"mail"`
	SMS      SMSConfig      `yaml:"sms"`
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
	PoolSize uint16 `yaml:"pool_size"`
}

type MailConfig struct {
	FromEmail    string `yaml:"from_email"`
	FromPassword string `yaml:"from_password"`
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
}

type SMSConfig struct {
	MessagingServiceSID string `yaml:"messaging_service_sid"`
	AccountSID          string `yaml:"account_sid"`
	AuthToken           string `yaml:"authentication_token"`
}
