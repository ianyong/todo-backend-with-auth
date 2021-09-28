package dbrepositories

import (
	"github.com/jmoiron/sqlx"

	"github.com/ianyong/todo-backend/internal/core/domainmodels"
)

type TodoDatabaseRepository struct {
	db *sqlx.DB
}

func (r *TodoDatabaseRepository) Get(id int64) (*domainmodels.Todo, error) {
	var todo domainmodels.Todo
	row := r.db.QueryRow("SELECT * FROM todos WHERE id = $1", id)
	err := row.Scan(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoDatabaseRepository) Add(todo *domainmodels.Todo) error {
	_, err := r.db.Exec(
		"INSERT INTO todos (name, description, due_date, is_completed) VALUES ($1, $2, $3, $4)",
		todo.GetName(),
		todo.GetDescription(),
		todo.GetDueDate(),
		todo.GetIsCompleted(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoDatabaseRepository) Update(todo *domainmodels.Todo) error {
	_, err := r.db.Exec(
		"UPDATE todos SET name = $1, description = $2, due_date = $3, is_completed = $4 WHERE id = $5",
		todo.GetName(),
		todo.GetDescription(),
		todo.GetDueDate(),
		todo.GetIsCompleted(),
		todo.GetID(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoDatabaseRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}