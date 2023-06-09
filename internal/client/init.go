package client

import (
	"database/sql"
	"notification-service/internal/config"
	"notification-service/internal/util"
)

var (
	Database *sql.DB
	RabbitMQ *RabbitMQClient
	Mail     IMail
	SMS      ISMS
	Push     IPush
)

func InitClients() {
	Database = InitDatabaseClient(&config.Master.Database)

	if config.Master.Service.APIs.Has(config.RABBITMQ_V1_API) {
		RabbitMQ = InitRabbitMQClient(&config.Master.RabbitMQ)
	}

	Mail = InitMailClient(
		&config.Master.Mail,
		config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_MAIL),
	)

	SMS = InitSMSClient(
		&config.Master.SMS,
		config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_SMS),
	)

	Push = InitPushClient(
		config.FirebaseAdminSDKPath,
		config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_PUSH),
	)

	util.Logger.Info().Msg("Loaded all clients")
}
