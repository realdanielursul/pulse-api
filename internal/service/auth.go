package service

import (
	"context"
	"fmt"
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
	_, err := s.userRepo.GetUserByLoginAndPassword(ctx, input.Login, s.passwordHasher.Hash(input.Password))
	if err != nil {
		return "", ErrInvalidLoginOrPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: input.Login,
	})

	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {
		return "", ErrCannotSignToken
	}

	if err := s.tokenRepo.CreateToken(ctx, &entity.Token{Login: input.Login, TokenString: tokenString, IsValid: true}); err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, tokenString string) (string, bool, error) {
	_, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signKey), nil
	})

	if err != nil {
		return "", false, ErrCannotParseToken
	}

	token, err := s.tokenRepo.GetToken(ctx, tokenString)
	if err != nil {
		return "", false, err
	}

	return token.Login, token.IsValid, nil
}

func (s *AuthService) UpdatePassword(ctx context.Context, userLogin string, input *AuthUpdatePasswordInput) error {
	_, err := s.userRepo.GetUserByLoginAndPassword(ctx, userLogin, s.passwordHasher.Hash(input.OldPassword))
	if err != nil {
		return ErrInvalidLoginOrPassword
	}

	if err := s.userRepo.UpdatePassword(ctx, userLogin, s.passwordHasher.Hash(input.NewPassword)); err != nil {
		return err
	}

	if err := s.tokenRepo.InvalidateUserTokens(ctx, userLogin); err != nil {
		return err
	}

	return nil
}
