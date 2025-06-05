package repository

import "github.com/example/todo/internal/domain"

// MemoryRepository is an in-memory implementation of TodoRepository.
type MemoryRepository struct {
	data   map[int64]*domain.Todo
	nextID int64
}

// NewMemoryRepository creates a MemoryRepository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: make(map[int64]*domain.Todo), nextID: 1}
}

func (m *MemoryRepository) Create(todo *domain.Todo) error {
	todo.ID = m.nextID
	m.nextID++
	m.data[todo.ID] = todo
	return nil
}

func (m *MemoryRepository) Get(id int64) (*domain.Todo, error) {
	if todo, ok := m.data[id]; ok {
		return todo, nil
	}
	return nil, domain.ErrNotFound
}

func (m *MemoryRepository) List() ([]*domain.Todo, error) {
	todos := make([]*domain.Todo, 0, len(m.data))
	for _, t := range m.data {
		todos = append(todos, t)
	}
	return todos, nil
}

func (m *MemoryRepository) Update(todo *domain.Todo) error {
	if _, ok := m.data[todo.ID]; !ok {
		return domain.ErrNotFound
	}
	m.data[todo.ID] = todo
	return nil
}

func (m *MemoryRepository) Delete(id int64) error {
	delete(m.data, id)
	return nil
}
