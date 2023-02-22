package todo

import (
	"database/sql"
	domainhttp "detroit-testing/app/domain/todo/http"
	"detroit-testing/app/domain/todo/repository"
	"detroit-testing/app/domain/todo/usecase"

	"github.com/gin-gonic/gin"
)

type BaseTodoServer struct {
	router      *gin.RouterGroup
	todoHandler domainhttp.TodoHandler
}

func NewBaseTodoServer(router *gin.RouterGroup, db *sql.DB) BaseTodoServer {
	todoRepo := repository.NewTodoRepo(db)
	todoSvc := usecase.NewTodoSvc(todoRepo)
	todoHandler := domainhttp.NewTodoHandler(todoSvc)
	return BaseTodoServer{
		router:      router,
		todoHandler: todoHandler,
	}
}

func (b BaseTodoServer) Build() {
	todo := b.router.Group("todos")
	todo.GET("/health", b.todoHandler.Health)
	todo.POST("", b.todoHandler.Create)
}
