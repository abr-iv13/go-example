package offers

import (
	"project/back/internal/repository"
)

type Service struct {
	repo repository.OfferRepo
}

func New(repo repository.OfferRepo) *Service {
	return &Service{repo: repo}
}
