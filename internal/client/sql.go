package client

import (
	"database/sql"
	"fmt"
	"notification-service/internal/config"
	"notification-service/internal/util"
	"time"
)

func InitDatabaseClient(conf *config.DatabaseConfig) *sql.DB {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Username, conf.Password, conf.Host, conf.Name)

	c, err := sql.Open(conf.Driver, connectionStr)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("Failed to initialize database client")
	}

	c.SetConnMaxIdleTime(5 * time.Second)
	c.SetConnMaxLifetime(0)
	c.SetMaxIdleConns(conf.PoolSize)
	c.SetMaxOpenConns(conf.PoolSize)

	return c
}
