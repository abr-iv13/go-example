package services

import (
	"project/back/internal/repository"
	"project/back/internal/services/offers"
	"project/back/internal/services/users"
)

//Service
type Users interface {
	Create()
}

type Offers interface{}

type Services struct {
	Users  Users
	Offers Offers
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Users:  users.New(repos.Users),
		Offers: offers.New(repos.Offers),
	}
}
