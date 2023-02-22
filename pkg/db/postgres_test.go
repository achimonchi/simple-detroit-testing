package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB(func() DBConfig {
		return DBConfig{
			Host:        "db",
			Port:        "5432",
			User:        "noobee",
			DBName:      "blog",
			Pass:        "iniPassword",
			MaxOpenConn: 10,
		}
	})
	fmt.Println(err.Error())

	require.Nil(t, err)
	require.NotNil(t, db)
}
