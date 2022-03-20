package services

import (
	"project/back/internal/contract"
	"project/back/internal/repository"
	"project/back/internal/services/offers"
	"project/back/internal/services/users"
)

type Services struct {
	Users  contract.UsersService
	Offers contract.OffersService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Users:  users.New(repos.UsersRepo),
		Offers: offers.New(repos.OffersRepo),
	}
}
