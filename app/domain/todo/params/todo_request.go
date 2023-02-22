package params

import (
	"detroit-testing/app/domain/todo/model"
	"time"
)

type CreateTodoRequest struct {
	Name        string
	Description string
}

func (c CreateTodoRequest) ToModel() model.Todo {
	return model.Todo{
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   time.Now(),
	}
}
