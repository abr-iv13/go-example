package services

import "project/back/internal/repository"

type OfferServiceInterface interface {
	Create()
}

type OfferService struct {
	repo repository.OfferRepoInterface
}

func NewOfferService(repo repository.OfferRepoInterface) *OfferService {
	return &OfferService{repo: repo}
}

func (s *OfferService) Create() {

}
