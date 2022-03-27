package users

import (
	"project/back/internal/repository"
)

type Service struct {
	repo repository.UserRepo
}

func New(repo repository.UserRepo) *Service {
	return &Service{repo: repo}
}
