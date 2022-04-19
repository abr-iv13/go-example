package services

import (
	"project/back/internal/repository"
)

//Service

type Services struct {
	UserService  UserServiceInterface
	OfferService OfferServiceInterface
}

func NewServices(repos *repository.Repository) *Services {
	userService := NewUserService(repos.UserRepo)
	offerService := NewOfferService(repos.OfferRepo)

	return &Services{
		UserService:  userService,
		OfferService: offerService,
	}
}
