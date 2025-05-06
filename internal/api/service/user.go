package service

import (
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/model"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}
