package sqlite

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/example/todo/internal/domain"
)

// TodoRepository implements domain.TodoRepository using Bun and SQLite.
type TodoRepository struct {
	DB *bun.DB
}

// NewTodoRepository creates a TodoRepository.
func NewTodoRepository(db *bun.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	_, err := r.DB.NewInsert().Model(todo).Exec(context.Background())
	return err
}

func (r *TodoRepository) Get(id int64) (*domain.Todo, error) {
	todo := new(domain.Todo)
	err := r.DB.NewSelect().Model(todo).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) List() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	err := r.DB.NewSelect().Model(&todos).Order("id").Scan(context.Background())
	return todos, err
}

func (r *TodoRepository) Update(todo *domain.Todo) error {
	_, err := r.DB.NewUpdate().Model(todo).WherePK().Exec(context.Background())
	return err
}

func (r *TodoRepository) Delete(id int64) error {
	_, err := r.DB.NewDelete().Model((*domain.Todo)(nil)).Where("id = ?", id).Exec(context.Background())
	return err
}
