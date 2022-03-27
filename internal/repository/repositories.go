package repository

import (
	"project/back/internal/repository/psql"
	"project/back/pkg/postgres"
)

//Repository
type UserRepo interface {
	Create()
	Get()
	Update()
	Delete()
}
type OfferRepo interface {
	Create()
	Get()
	Update()
	Delete()
}

type Repositories struct {
	UserRepo  UserRepo
	OfferRepo OfferRepo
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		UserRepo:  psql.NewUser(pg),
		OfferRepo: psql.NewOffer(pg),
	}
}
