package service

import (
	"backend_food_delivery/model"
	"backend_food_delivery/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "asdddqwwee"
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePassHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
