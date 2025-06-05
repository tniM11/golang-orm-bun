package domain

import "time"

// Todo represents a task item.
type Todo struct {
	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
}

// TodoRepository defines data access methods for Todo.
type TodoRepository interface {
	Create(todo *Todo) error
	Get(id int64) (*Todo, error)
	List() ([]*Todo, error)
	Update(todo *Todo) error
	Delete(id int64) error
}
