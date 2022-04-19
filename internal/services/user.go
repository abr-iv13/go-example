package services

import (
	"project/back/internal/repository"
)

type UserServiceInterface interface {
	Create()
}

type UserService struct {
	repo repository.UserRepoInterface
}

func NewUserService(repo repository.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create() {
	s.repo.Create()
}
