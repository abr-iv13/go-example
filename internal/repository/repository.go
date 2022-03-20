package repository

import (
	"project/back/internal/contract"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UsersRepo  contract.UserRepo
	OffersRepo contract.OffersRepo
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UsersRepo:  NewUserRepo(db),
		OffersRepo: NewOfferRepo(db),
	}

}
