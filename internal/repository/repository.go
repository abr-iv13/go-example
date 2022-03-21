package repository

import (
	"project/back/internal/repository/psql"

	"github.com/jmoiron/sqlx"
)

//Repository
type Users interface {
	Create()
	Get()
	Update()
	Delete()
}
type Offers interface {
	Create()
	Get()
	Update()
	Delete()
}

type Repositories struct {
	Users  Users
	Offers Offers
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users:  psql.NewUsers(db),
		Offers: psql.NewOffers(db),
	}
}
