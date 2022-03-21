package psql

import (
	"github.com/jmoiron/sqlx"
)

type Offers struct {
	db *sqlx.DB
}

func NewOffers(db *sqlx.DB) *Offers {
	return &Offers{db: db}
}

func (r *Offers) Create() {
}

func (r *Offers) Get() {
}

func (r *Offers) Update() {
}

func (r *Offers) Delete() {
}
