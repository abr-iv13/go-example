package repository

import (
	"project/back/pkg/postgres"
)

type OfferRepoInterface interface {
	Create()
	Get()
	Update()
	Delete()
}
type UserPostgres struct {
	db *postgres.Postgres
}

func NewUserPostgres(pg *postgres.Postgres) *UserPostgres {
	return &UserPostgres{db: pg}
}

func (u *UserPostgres) Create() {

}

func (u *UserPostgres) Get() {

}

func (u *UserPostgres) Update() {

}

func (u *UserPostgres) Delete() {

}
