package repository

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, email, password string) (domain.User, error)
}
type Repository struct {
	Authorization
}

func NewRepository (db *mongo.Database) *Repository{
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
