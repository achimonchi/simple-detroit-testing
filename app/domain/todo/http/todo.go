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
			"message": "BAD REQUEST",
			"error":   err.Error(),
		})
		return
	}

	newCtx := ctx.Request.Context()
	if err := t.todoSvc.CreateTodo(newCtx, req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "CREATE SUCCESS",
	})

}

func (t TodoHandler) GetAll(ctx *gin.Context) {
	newCtx := ctx.Request.Context()

	resp, err := t.todoSvc.GetAllTodos(newCtx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
			"error":   err.Error(),
		})
		return
	}

	if !resp.IsNoError {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "INTERNAL SERVER ERROR",
			"error":   "unknown error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET SUCCESS",
		"payload": resp.Todos,
	})
}
