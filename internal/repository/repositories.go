package repository

import (
	"project/back/internal/repository/psql"
	"project/back/pkg/postgres"
)

//Repository

type Repositories struct {
	UserRepo  psql.UserRepoInterface
	OfferRepo psql.OfferRepoInterface
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		UserRepo:  psql.NewUserRepo(pg),
		OfferRepo: psql.NewOfferRepo(pg),
	}
}
