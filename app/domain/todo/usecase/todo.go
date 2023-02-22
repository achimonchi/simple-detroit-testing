package usecase

import (
	"context"
	"detroit-testing/app/domain/todo/params"
)

type TodoSvc struct {
	todoRepo TodoRepo
}

func NewTodoSvc(todoRepo TodoRepo) TodoSvc {
	return TodoSvc{
		todoRepo: todoRepo,
	}
}

func (t TodoSvc) CreateTodo(ctx context.Context, req params.CreateTodoRequest) error {
	todo := req.ToModel()
	err := t.todoRepo.CreateTodo(ctx, todo)
	return err
}
