package service

import (
	"backend_food_delivery/pkg/domain"
	"backend_food_delivery/pkg/repository"
	"context"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService{
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(ctx context.Context, list domain.Todo)  error {
	return s.repo.Create(ctx, list)
}

func (s *TodoListService) GetAll(ctx context.Context) ([]domain.Todo, error){
	return s.repo.GetAll(ctx)
}