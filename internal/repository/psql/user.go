package psql

import (
	"project/back/pkg/postgres"
)

type OfferRepoInterface interface {
	Create()
	Get()
	Update()
	Delete()
}
type UserRepo struct {
	db *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{db: pg}
}

func (u *UserRepo) Create() {

}

func (u *UserRepo) Get() {

}

func (u *UserRepo) Update() {

}

func (u *UserRepo) Delete() {

}
