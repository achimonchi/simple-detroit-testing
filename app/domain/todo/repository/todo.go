package repository

import (
	"context"
	"database/sql"
	"detroit-testing/app/domain/todo/model"
)

type todoRepo struct {
	db *sql.DB
}

// CreateTodo implements usecase.TodoRepo
func (todoRepo) CreateTodo(ctx context.Context, todo model.Todo) error {
	panic("unimplemented")
}

func NewTodoRepo(db *sql.DB) todoRepo {
	return todoRepo{
		db: db,
	}
}
