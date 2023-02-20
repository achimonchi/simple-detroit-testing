package todo

import (
	domainhttp "detroit-testing/app/domain/todo/http"

	"github.com/gin-gonic/gin"
)

type BaseTodoServer struct {
	router      *gin.RouterGroup
	todoHandler domainhttp.TodoHandler
}

func NewBaseTodoServer(router *gin.RouterGroup) BaseTodoServer {
	todoHandler := domainhttp.NewTodoHandler()
	return BaseTodoServer{
		router:      router,
		todoHandler: todoHandler,
	}
}

func (b BaseTodoServer) Build() {
	todo := b.router.Group("todos")
	todo.GET("/health", b.todoHandler.Health)
}
