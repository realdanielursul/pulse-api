package service

import (
	"context"

	"github.com/realdanielursul/pulse-api/internal/repository"
)

type UserService struct {
	userRepo    repository.User
	friendRepo  repository.Friend
	countryRepo repository.Country
}

func NewUserService(userRepo repository.User, friendRepo repository.Friend, countryRepo repository.Country) *UserService {
	return &UserService{
		userRepo:    userRepo,
		friendRepo:  friendRepo,
		countryRepo: countryRepo,
	}
}

func (s *UserService) GetProfile(ctx context.Context, login string, requesterLogin string) (*UserOutput, error) {
	dbUser, err := s.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		if dbUser == nil {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	user := &UserOutput{
		Login:       dbUser.Login,
		Email:       dbUser.Email,
		CountryCode: dbUser.CountryCode,
		IsPublic:    dbUser.IsPublic,
		Phone:       dbUser.Phone,
		Image:       dbUser.Image,
	}

	if login == requesterLogin {
		return user, nil
	}

	if user.IsPublic {
		return user, nil
	}

	isFriend, err := s.friendRepo.IsFriend(ctx, requesterLogin, login)
	if err != nil {
		return nil, err
	}

	if isFriend {
		return user, nil
	}

	return nil, ErrAccessDenied
}

func (s *UserService) GetMyProfile(ctx context.Context, userLogin string) (*UserOutput, error) {
	user, err := s.userRepo.GetUserByLogin(ctx, userLogin)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		Login:       user.Login,
		Email:       user.Email,
		CountryCode: user.CountryCode,
		IsPublic:    user.IsPublic,
		Phone:       user.Phone,
		Image:       user.Image,
	}, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userLogin string, input *UserUpdateProfileInput) (*UserOutput, error) {
	country, err := s.countryRepo.GetCountryByAlpha2(ctx, *input.CountryCode)
	if err != nil {
		if country == nil {
			return nil, ErrCountryNotFound
		}

		return nil, err
	}

	if input.Phone != nil {
		if _, err := s.userRepo.GetUserByPhone(ctx, *input.Phone); err == nil {
			return nil, ErrPhoneAlreadyExists
		}
	}

	if err := s.userRepo.UpdateUser(ctx, userLogin, input.CountryCode, input.Phone, input.Image, input.IsPublic); err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByLogin(ctx, userLogin)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		Login:       user.Login,
		Email:       user.Email,
		CountryCode: user.CountryCode,
		IsPublic:    user.IsPublic,
		Phone:       user.Phone,
		Image:       user.Image,
	}, nil
}
