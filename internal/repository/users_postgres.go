package repository

import (
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	*sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create() {
}

func (u *UserRepo) Get() {
}

func (u *UserRepo) Update() {
}

func (u *UserRepo) Delete() {
}
