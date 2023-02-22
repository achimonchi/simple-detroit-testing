package usecase

import (
	"context"
	"detroit-testing/app/domain/todo/model"
)

type TodoRepo interface {
	TodoRepoRead
	TodoRepoWrite
}

type TodoRepoRead interface {
	TodoRepoReadAll
}

type TodoRepoReadAll interface {
	GetAll(ctx context.Context) ([]model.Todo, error)
}

type TodoRepoWrite interface {
	CreateTodo(ctx context.Context, todo model.Todo) error
}
