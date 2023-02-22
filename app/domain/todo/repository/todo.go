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

// GetAll implements usecase.TodoRepo
func (r todoRepo) GetAll(ctx context.Context) ([]model.Todo, error) {
	query := `
		SELECT id, name, description, created_at
		FROM todos
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos = []model.Todo{}

	for rows.Next() {
		var todo = model.Todo{}
		err := rows.Scan(
			&todo.Id,
			&todo.Name,
			&todo.Description,
			&todo.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
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
