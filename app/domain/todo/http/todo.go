package domainhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
}

func NewTodoHandler() TodoHandler {
	return TodoHandler{}
}

func (t TodoHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
