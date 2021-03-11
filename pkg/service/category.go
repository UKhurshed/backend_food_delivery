package service

import (
	"backend_food_delivery/pkg/domain"
	"backend_food_delivery/pkg/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, ctg domain.Category) error {
	return s.repo.CreateCategory(ctx, ctg)
}

func (s *CategoryService) GetAllCategory(ctx context.Context) ([]domain.Category, error){
	return s.repo.GetAllCategory(ctx)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id primitive.ObjectID) error{
	return s.repo.DeleteCategory(ctx, id)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, inp domain.Category) error{
	return s.repo.UpdateCategory(ctx, inp)
}
