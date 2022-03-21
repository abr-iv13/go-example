package offers

import (
	"project/back/internal/repository"
)

type Service struct {
	repo repository.Offers
}

func New(repo repository.Offers) *Service {
	return &Service{repo: repo}

}
