package params

import "detroit-testing/app/domain/todo/model"

type GetAllTodosResponse struct {
	IsNoError bool
	Todos     []model.Todo
}
