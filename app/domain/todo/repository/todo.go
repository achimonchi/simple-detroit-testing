package repository

import (
	"context"
	"database/sql"
	"detroit-testing/app/domain/todo/model"
)

func NewTodoRepo(db *sql.DB) todoRepo {
	return todoRepo{
		db: db,
	}
}

type todoRepo struct {
	db *sql.DB
}

// CreateTodo implements usecase.TodoRepo
func (r todoRepo) CreateTodo(ctx context.Context, todo model.Todo) error {
	query := `
		INSERT INTO todos (name, description, created_at)
		VALUES ($1, $2, $3)
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, todo.Name, todo.Description, todo.CreatedAt)
	return err
}
