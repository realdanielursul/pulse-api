package service

import (
	"time"

	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

type Repository interface {
}

type Service struct {
	repo           Repository
	passwordHasher hasher.SHA1Hasher
	signKey        string
	tokenTTL       time.Duration
}

func NewService(repo Repository, passwordHasher hasher.SHA1Hasher, signKey string, tokenTTL time.Duration, salt string) *Service {
	return &Service{
		repo:           repo,
		passwordHasher: passwordHasher,
		signKey:        signKey,
		tokenTTL:       tokenTTL,
	}
}
