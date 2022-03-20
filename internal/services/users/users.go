package users

import "project/back/internal/contract"

type Service struct {
	repo contract.UserRepo
}

func New(repo contract.UserRepo) *Service {
	return &Service{repo: repo}
}
