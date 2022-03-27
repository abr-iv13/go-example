package psql

import (
	"project/back/pkg/postgres"
)

type User struct {
	db *postgres.Postgres
}

func NewUser(pg *postgres.Postgres) *User {
	return &User{db: pg}
}

func (u *User) Create() {

}

func (u *User) Get() {

}

func (u *User) Update() {

}

func (u *User) Delete() {

}
