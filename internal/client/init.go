package client

import (
	"database/sql"
	"notification-service/internal/config"
)

var (
	Database *sql.DB
	Mail     IMail
	SMS      ISMS
	Push     IPush
)

func InitClients() {
	Database = InitDatabaseClient(&config.Master.Database)

	Mail = InitMailClient(
		&config.Master.Mail,
		!config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_MAIL),
	)

	SMS = InitSMSClient(
		&config.Master.SMS,
		!config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_SMS),
	)
}
