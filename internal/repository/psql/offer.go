package psql

import (
	"project/back/pkg/postgres"
)

type UserRepoInterface interface {
	Create()
	Get()
	Update()
	Delete()
}
type OfferRepo struct {
	db *postgres.Postgres
}

func NewOfferRepo(pg *postgres.Postgres) *OfferRepo {
	return &OfferRepo{db: pg}
}

func (r *OfferRepo) Create() {

}

func (r *OfferRepo) Get() {

}

func (r *OfferRepo) Update() {

}

func (r *OfferRepo) Delete() {

}
