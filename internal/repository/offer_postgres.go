package repository

import (
	"project/back/pkg/postgres"
)

type UserRepoInterface interface {
	Create()
	Get()
	Update()
	Delete()
}
type OfferPostgres struct {
	db *postgres.Postgres
}

func NewOfferPostgres(pg *postgres.Postgres) *OfferPostgres {
	return &OfferPostgres{db: pg}
}

func (r *OfferPostgres) Create() {

}

func (r *OfferPostgres) Get() {

}

func (r *OfferPostgres) Update() {

}

func (r *OfferPostgres) Delete() {

}
