package infra

import (
	"detroit-testing/app/domain/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.RouterGroup
}

func NewHandler(router *gin.RouterGroup) Handler {
	return Handler{
		router: router,
	}
}

func (h Handler) BuildTodo() todo.BaseTodoServer {
	return todo.NewBaseTodoServer(h.router)
}
