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

func (t TodoSvc) GetAllTodos(ctx context.Context) (params.GetAllTodosResponse, error) {
	var resp = params.GetAllTodosResponse{
		IsNoError: false,
	}
	todos, err := t.todoRepo.GetAll(ctx)
	if err != nil {
		return resp, err
	}

	resp.IsNoError = true
	resp.Todos = todos
	return resp, nil
}
