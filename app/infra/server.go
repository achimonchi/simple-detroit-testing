package infra

import (
	"github.com/gin-gonic/gin"
)

type Infra struct {
	port string
}

func NewInfra(port string) Infra {
	return Infra{
		port: port,
	}
}

func (i Infra) RunGin() {
	router := gin.New()

	v1 := router.Group("v1")

	handler := NewHandler(v1)
	handler.BuildTodo().Build()

	router.Run(i.port)
}
