package services

import (
	"project/back/internal/repository"
	"project/back/internal/services/offers"
	"project/back/internal/services/users"
)

//Service
type UserService interface {
	Create()
}

type OfferService interface {
	Create()
}

type Services struct {
	User  UserService
	Offer OfferService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User:  users.New(repos.UserRepo),
		Offer: offers.New(repos.OfferRepo),
	}
}
