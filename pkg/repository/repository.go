package repository

import (
	"backend_food_delivery/model"
	"github.com/go-pg/pg/v10"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *pg.DB) *Repository{
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}