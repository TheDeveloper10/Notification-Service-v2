package config

import "notification-service/internal/util"

type MasterConfig struct {
	Service  ServiceConfig  `yaml:"service"`
	HTTP     HTTPConfig     `yaml:"http"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
	Database DatabaseConfig `yaml:"database"`
	Mail     MailConfig     `yaml:"mail"`
	SMS      SMSConfig      `yaml:"sms"`
}

type ServiceConfig struct {
	APIs              util.Strings `yaml:"apis"`
	NotificationTypes util.Strings `yaml:"notification_types"`
	AllowedLanguages  util.Strings `yaml:"allowed_languages"`
	Auth              AuthConfig   `yaml:"auth"`
	Cache             CacheConfig  `yaml:"cache"`
}

type AuthConfig struct {
	MasterClientID     string `yaml:"master_client_id"`
	MasterClientSecret string `yaml:"master_client_secret"`
	TokenExpiryTime    uint32 `yaml:"token_expiry_time"`
}

type CacheConfig struct {
	TemplatesCacheEntryExpiry uint32 `yaml"templates_cache_entry_expiry"`
}

type HTTPConfig struct {
	Address string `yaml:"address"`
}

type RabbitMQConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
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
	FromPhoneNumber string `yaml:"from_phone_number"`
	AccountSID      string `yaml:"account_sid"`
	AuthToken       string `yaml:"authentication_token"`
}
