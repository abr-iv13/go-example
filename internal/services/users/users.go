package users

import (
	"project/back/internal/repository"
)

type Service struct {
	repo repository.Users
}

func New(repo repository.Users) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create() {
	s.repo.Create()
}
