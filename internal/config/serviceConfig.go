package config

import "notification-service/internal/util"

type MasterConfig struct {
	Service  ServiceConfig  `json:"service"`
	HTTP     HTTPConfig     `json:"http"`
	RabbitMQ RabbitMQConfig `json:"rabbitmq"`
	Database DatabaseConfig `json:"database"`
	Mail     MailConfig     `json:"mail"`
	SMS      SMSConfig      `json:"sms"`
}

type ServiceConfig struct {
	APIs              util.Strings `json:"apis"`
	NotificationTypes util.Strings `json:"notification_types"`
	AllowedLanguages  util.Strings `json:"allowed_languages"`
	Auth              AuthConfig   `json:"auth"`
	Cache             CacheConfig  `json:"cache"`
}

type AuthConfig struct {
	MasterClientID     string `json:"master_client_id"`
	MasterClientSecret string `json:"master_client_secret"`
	TokenExpiryTime    uint32 `json:"token_expiry_time"`
	MaxActiveClients   uint32 `json:"max_active_clients"`
}

type CacheConfig struct {
	TemplatesCacheLimit       uint32 `json:"templates_cache_limit"`
	TemplatesCacheEntryExpiry uint32 `json:"templates_cache_entry_expiry"`
	TemplatesCacheCleanupTime uint32 `json:"templates_cache_cleanup_time"`
}

type HTTPConfig struct {
	Address string `json:"address"`
}

type RabbitMQConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	PoolSize uint16 `json:"pool_size"`
}

type MailConfig struct {
	FromEmail    string `json:"from_email"`
	FromPassword string `json:"from_password"`
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
}

type SMSConfig struct {
	FromPhoneNumber string `json:"from_phone_number"`
	AccountSID      string `json:"account_sid"`
	AuthToken       string `json:"authentication_token"`
}
