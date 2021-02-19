package service

import (
	"backend_food_delivery/model"
	"backend_food_delivery/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
