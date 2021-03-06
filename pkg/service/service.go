package service

import (
	"backend_food_delivery/pkg/domain"
	"backend_food_delivery/pkg/repository"
	"context"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) error
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(token string) (string, error)
}

type TodoList interface {
	Create(ctx context.Context, list domain.Todo) error
	GetAll(ctx context.Context) ([]domain.Todo, error)
}

type Category interface {
	CreateCategory(ctx context.Context, category domain.Category) error
	GetAllCategory(ctx context.Context) ([]domain.Category, error)
}

type Service struct {
	Authorization
	TodoList
	Category
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		Category:      NewCategoryService(repos.Category),
	}
}
