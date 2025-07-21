package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) Register(req models.RegisterRequest) (models.UserProfile, error) {
	if err := validateLogin(s, req.Login); err != nil {
		return models.UserProfile{}, err
	}

	if err := validateEmail(s, req.Email); err != nil {
		return models.UserProfile{}, err
	}

	if err := validatePassword(req.Password); err != nil {
		return models.UserProfile{}, err
	}

	if err := validateCountryCode(s, req.CountryCode); err != nil {
		return models.UserProfile{}, err
	}

	if err := validatePhone(s, req.Phone); err != nil {
		return models.UserProfile{}, err
	}

	if err := validateImage(req.Image); err != nil {
		return models.UserProfile{}, err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.Register(req)
}

func (s *Service) SignIn(req models.SignInRequest) (string, error) {
	req.Password = generatePasswordHash(req.Password)

	login, err := s.repo.SignIn(req)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", errors.ErrInvalidUsernameOrPassword
		}

		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: login,
	})

	token, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	if err := s.repo.AddToken(login, token); err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) ValidateToken(token string) error {
	return s.repo.ValidateToken(token)
}
