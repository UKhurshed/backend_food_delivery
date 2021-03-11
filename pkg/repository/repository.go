package repository

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, email, password string) (domain.User, error)
}
type TodoList interface {
	Create(ctx context.Context, list domain.Todo) error
	GetAll(ctx context.Context) ([]domain.Todo, error)
}

type Category interface {
	CreateCategory(ctx context.Context, ctg domain.Category) error
	GetAllCategory(ctx context.Context) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, id primitive.ObjectID) error
	UpdateCategory(ctx context.Context, inp domain.Category) error
}

type Repository struct {
	Authorization
	TodoList
	Category
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		TodoList:      NewTodoListMongo(db),
		Category: NewCategoryMongo(db),
	}
}
