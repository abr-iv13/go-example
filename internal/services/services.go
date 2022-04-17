package services

import (
	"project/back/internal/repository"
	"project/back/internal/services/offers"
	"project/back/internal/services/users"
)

//Service

type Services struct {
	UserService  users.UserServiceInterface
	OfferService offers.OfferServiceInterface
}

func NewServices(repos *repository.Repositories) *Services {
	userService := users.NewUserService(repos.UserRepo)
	offerService := offers.NewOfferService(repos.OfferRepo)

	return &Services{
		UserService:  userService,
		OfferService: offerService,
	}
}
