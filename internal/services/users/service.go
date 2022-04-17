package users

import (
	"project/back/internal/repository/psql"
)

type UserService struct {
	repo psql.UserRepoInterface
}

func NewUserService(repo psql.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}
