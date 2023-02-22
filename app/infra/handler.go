package infra

import (
	"database/sql"
	"detroit-testing/app/domain/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.RouterGroup
	db     *sql.DB
}

func NewHandler(router *gin.RouterGroup) *Handler {
	return &Handler{
		router: router,
	}
}

func (h *Handler) SetDb(db *sql.DB) *Handler {
	h.db = db
	return h
}

func (h *Handler) BuildTodo() todo.BaseTodoServer {
	return todo.NewBaseTodoServer(h.router, h.db)
}
