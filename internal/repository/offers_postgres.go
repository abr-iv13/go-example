package repository

import (
	"github.com/jmoiron/sqlx"
)

type OfferRepo struct {
	*sqlx.DB
}

func NewOfferRepo(db *sqlx.DB) *OfferRepo {

	return &OfferRepo{db}
}

func (r *OfferRepo) Create() {
}

func (r *OfferRepo) Get() {
}

func (r *OfferRepo) Update() {
}

func (r *OfferRepo) Delete() {
}
