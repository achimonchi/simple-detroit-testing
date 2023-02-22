package usecase

import (
	"context"
	"detroit-testing/app/domain/todo/model"
)

type TodoRepo interface {
	TodoRepoRead
	TodoRepoWrite
}

type TodoRepoRead interface{}
type TodoRepoWrite interface {
	CreateTodo(ctx context.Context, todo model.Todo) error
}
