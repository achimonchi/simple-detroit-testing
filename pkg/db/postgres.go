package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host        string
	Port        string
	User        string
	Pass        string
	DBName      string
	MaxOpenConn int
}

func ConnectDB(conf func() DBConfig) (*sql.DB, error) {
	dbConf := conf()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.User, dbConf.Pass, dbConf.DBName)

	db, err := sql.Open("postgres", dsn)
	fmt.Sprintln(db)
	if err != nil {
		fmt.Sprintln("error:", err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(dbConf.MaxOpenConn)

	if err := db.Ping(); err != nil {
		fmt.Sprintln("error:", err.Error())
		return nil, err
	}

	return db, nil
}
