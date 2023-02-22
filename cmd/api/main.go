package main

import (
	"detroit-testing/app/infra"
	"detroit-testing/config"
	"detroit-testing/pkg/db"
	"log"
)

func main() {
	// init config
	err := config.LoadConfig("./.env")
	if err != nil {
		log.Println("Skipped load config from env file")
	}

	db, err := db.ConnectDB(func() db.DBConfig {
		return db.DBConfig{
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

	if err != nil {
		panic(err)
	}

	server := infra.NewInfra(":9999", db)
	server.RunGin()
}
