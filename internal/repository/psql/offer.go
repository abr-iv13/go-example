package psql

import (
	"project/back/pkg/postgres"
)

type Offer struct {
	db *postgres.Postgres
}

func NewOffer(pg *postgres.Postgres) *Offer {
	return &Offer{db: pg}
}

func (r *Offer) Create() {

}

func (r *Offer) Get() {

}

func (r *Offer) Update() {

}

func (r *Offer) Delete() {

}
