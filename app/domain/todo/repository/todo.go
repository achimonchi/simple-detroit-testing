package repository

import "database/sql"

type todoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) todoRepo {
	return todoRepo{
		db: db,
	}
}
