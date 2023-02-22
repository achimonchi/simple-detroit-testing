package db

import (
	"detroit-testing/config"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	err := config.LoadConfig("./../../.testing.env")
	if err != nil {
		log.Println("Skipped load env config")
	}
}

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB(func() DBConfig {
		return DBConfig{
			Host:            config.GetString(config.DB_HOST, "localhost"),
			Port:            config.GetString(config.DB_PORT, "5432"),
			User:            config.GetString(config.DB_USER, "user"),
			Pass:            config.GetString(config.DB_PASS, "pass"),
			DBName:          config.GetString(config.DB_NAME, "db"),
			MaxOpenConn:     int(config.GetUint8(config.DB_MAX_OPEN_CONN, 10)),
			MaxIdleConn:     int(config.GetUint8(config.DB_MAX_IDLE_CONN, 10)),
			MaxLifetimeConn: int(config.GetUint8(config.DB_MAX_LIFETIME_CONN, 10)),
			MaxIdletimeConn: int(config.GetUint8(config.DB_MAX_IDLETIME_CONN, 10)),
		}
	})

	require.Nil(t, err)
	require.NotNil(t, db)
}
