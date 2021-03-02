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
type Service struct {
	Authorization
}

func NewService (repos *repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
