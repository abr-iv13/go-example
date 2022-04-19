package repository

import (
	"project/back/pkg/postgres"
)

//Repository

type Repository struct {
	UserRepo  UserRepoInterface
	OfferRepo OfferRepoInterface
}

func NewRepository(pg *postgres.Postgres) *Repository{
	return &Repository{
		UserRepo:  NewUserPostgres(pg),
		OfferRepo: NewOfferPostgres(pg),
	}
}
