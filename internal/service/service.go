package service

import (
	"time"

	"github.com/realdanielursul/pulse-api/internal/model"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

type Repository interface {
	ListCountries(regions []string) ([]model.Country, error)
	GetCountryByAlpha2(alpha2 string) (model.Country, error)
}

type Service struct {
	repo           Repository
	passwordHasher hasher.SHA1Hasher
	signKey        string
	tokenTTL       time.Duration
}

func NewService(repo Repository, passwordHasher hasher.SHA1Hasher, signKey string, tokenTTL time.Duration) *Service {
	return &Service{
		repo:           repo,
		passwordHasher: passwordHasher,
		signKey:        signKey,
		tokenTTL:       tokenTTL,
	}
}
