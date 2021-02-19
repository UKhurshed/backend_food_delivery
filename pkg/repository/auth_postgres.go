package repository

import (
	"backend_food_delivery/model"
	"github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
)

type AuthPostgres struct {
	db *pg.DB
}

func NewAuthPostgres(db *pg.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	return 0, nil
}
