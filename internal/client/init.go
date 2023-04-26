package client

import (
	"database/sql"
	"notification-service/internal/config"
)

var (
	Database *sql.DB
)

func InitClients() {
	Database = InitDatabaseClient(&config.Master.Database)
}
