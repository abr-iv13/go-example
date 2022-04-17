package offers

import (
	"project/back/internal/repository/psql"
)

type OfferService struct {
	repo psql.OfferRepoInterface
}

func NewOfferService(repo psql.OfferRepoInterface) *OfferService {
	return &OfferService{repo: repo}
}
