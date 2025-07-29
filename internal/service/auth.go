package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/realdanielursul/pulse-api/internal/entity"
	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

type TokenClaims struct {
	jwt.StandardClaims
	Login string
}

type AuthService struct {
	userRepo       repository.User
	tokenRepo      repository.Token
	passwordHasher hasher.PasswordHasher
	signKey        string
	tokenTTL       time.Duration
}

func NewAuthService(userRepo repository.User, tokenRepo repository.Token, passwordHasher hasher.PasswordHasher, signKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		tokenRepo:      tokenRepo,
		passwordHasher: passwordHasher,
		signKey:        signKey,
		tokenTTL:       tokenTTL,
	}
}

func (s *AuthService) Register(ctx context.Context, input *AuthRegisterInput) (*AuthRegisterOutput, error) {
	if _, err := s.userRepo.GetUserByLogin(ctx, input.Login); err == nil {
		return nil, ErrLoginAlreadyExists
	}

	if _, err := s.userRepo.GetUserByEmail(ctx, input.Email); err == nil {
		return nil, ErrEmailAlreadyExists
	}

	if _, err := s.userRepo.GetUserByPhone(ctx, input.Phone); err == nil {
		return nil, ErrPhoneAlreadyExists
	}

	user := &entity.User{
		Login:        input.Login,
		Email:        input.Email,
		PasswordHash: s.passwordHasher.Hash(input.Password),
		CountryCode:  input.CountryCode,
		IsPublic:     input.IsPublic,
		Phone:        input.Phone,
		Image:        input.Image,
	}

	newUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &AuthRegisterOutput{
		Login:       newUser.Login,
		Email:       newUser.Email,
		CountryCode: newUser.CountryCode,
		IsPublic:    newUser.IsPublic,
		Phone:       newUser.Phone,
		Image:       newUser.Image,
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, input *AuthSignInInput) (string, error) {
	panic("")
}

func (s *AuthService) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	panic("")
}

func (s *AuthService) UpdatePassword(ctx context.Context, userLogin string, input *AuthUpdatePasswordInput) error {
	panic("")
}
