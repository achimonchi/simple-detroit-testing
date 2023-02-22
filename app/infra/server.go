package infra

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Infra struct {
	port string
	db   *sql.DB
}

func NewInfra(port string, db *sql.DB) Infra {
	return Infra{
		port: port,
		db:   db,
	}
}

func (i Infra) RunGin() {
	router := gin.New()

	v1 := router.Group("v1")

	handler := NewHandler(v1)
	handler.BuildTodo().Build()

	router.Run(i.port)
}
