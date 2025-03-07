package services

import (
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/models"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/repositories"
)

type IUserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.Create(user)
}
