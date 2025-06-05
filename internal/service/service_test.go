package service

import (
	"testing"

	"github.com/example/todo/internal/domain"
	"github.com/example/todo/internal/repository"
)

func TestTodoService(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := NewTodoService(repo)

	todo := &domain.Todo{Title: "test"}
	if err := svc.Create(todo); err != nil {
		t.Fatalf("create: %v", err)
	}
	if todo.ID == 0 {
		t.Fatalf("expected ID set")
	}

	got, err := svc.Get(todo.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Title != todo.Title {
		t.Fatalf("got %q want %q", got.Title, todo.Title)
	}

	todo.Title = "changed"
	if err := svc.Update(todo); err != nil {
		t.Fatalf("update: %v", err)
	}

	list, err := svc.List()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected list length 1 got %d", len(list))
	}

	if err := svc.Delete(todo.ID); err != nil {
		t.Fatalf("delete: %v", err)
	}
	list, _ = svc.List()
	if len(list) != 0 {
		t.Fatalf("expected empty list")
	}
}
