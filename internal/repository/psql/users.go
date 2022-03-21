package psql

import (
	"github.com/jmoiron/sqlx"
)

type Users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) *Users {
	return &Users{db}
}

func (u *Users) Create() {
}

func (u *Users) Get() {
}

func (u *Users) Update() {
}

func (u *Users) Delete() {
}
