package domainhttp

import (
	"detroit-testing/app/domain/todo/repository"
	"detroit-testing/app/domain/todo/usecase"
	"detroit-testing/config"
	"detroit-testing/pkg/db"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var basePath = "/v1/todos"

var handler TodoHandler

func init() {
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

	repo := repository.NewTodoRepo(db)
	svc := usecase.NewTodoSvc(repo)
	handler = NewTodoHandler(svc)
}

func TestTodoHealthCheck(t *testing.T) {
	path := "/health"
	req, err := http.NewRequest("GET", basePath+path, nil)
	require.Nil(t, err)

	router := gin.Default()

	router.GET(basePath+path, handler.Health)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	fmt.Printf("%+v\n", rr)
	fmt.Printf("Body %v\n", rr.Body.String())

	require.Equal(t, 200, rr.Code)
}
