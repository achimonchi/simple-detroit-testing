package domainhttp

import (
	"detroit-testing/app/domain/todo/params"
	"detroit-testing/app/domain/todo/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoSvc usecase.TodoSvc
}

func NewTodoHandler(todoSvc usecase.TodoSvc) TodoHandler {
	return TodoHandler{
		todoSvc: todoSvc,
	}
}

func (t TodoHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (t TodoHandler) Create(ctx *gin.Context) {
	var req = params.CreateTodoRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "BAD_REQUEST",
			"error":   err.Error(),
		})
		return
	}
}
