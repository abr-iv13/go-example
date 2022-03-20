package offers

import (
	"project/back/internal/contract"
)

type Service struct {
	repo contract.OffersRepo
}

func New(repo contract.OffersRepo) *Service {
	return &Service{repo: repo}

}
