package service

import "github.com/example/todo/internal/domain"

// TodoService provides Todo use cases.
type TodoService struct {
	Repo domain.TodoRepository
}

// NewTodoService creates a new TodoService.
func NewTodoService(r domain.TodoRepository) *TodoService {
	return &TodoService{Repo: r}
}

func (s *TodoService) Create(todo *domain.Todo) error {
	return s.Repo.Create(todo)
}

func (s *TodoService) Get(id int64) (*domain.Todo, error) {
	return s.Repo.Get(id)
}

func (s *TodoService) List() ([]*domain.Todo, error) {
	return s.Repo.List()
}

func (s *TodoService) Update(todo *domain.Todo) error {
	return s.Repo.Update(todo)
}

func (s *TodoService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
