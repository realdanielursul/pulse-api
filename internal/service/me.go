package service

import (
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (s *Service) GetMyProfile(login string) (entity.UserProfile, error) {
	return s.repo.GetProfile(login)
}

func (s *Service) UpdateProfile(login string, req entity.UpdateProfileRequest) (entity.UserProfile, error) {
	if req.CountryCode != nil {
		if err := validateCountryCode(s, *req.CountryCode); err != nil {
			return entity.UserProfile{}, err
		}
	}

	if req.Phone != nil {
		if err := validatePhone(s, *req.Phone); err != nil {
			return entity.UserProfile{}, err
		}
	}

	if req.Image != nil {
		if err := validateImage(*req.Image); err != nil {
			return entity.UserProfile{}, err
		}
	}

	return s.repo.UpdateProfile(login, req)
}

func (s *Service) UpdatePassword(login string, req entity.UpdatePasswordRequest) error {
	if err := validatePassword(req.NewPassword); err != nil {
		return err
	}

	req.OldPassword = generatePasswordHash(req.OldPassword)
	req.NewPassword = generatePasswordHash(req.NewPassword)

	if err := s.repo.UpdatePassword(login, req); err != nil {
		return err
	}

	if err := s.repo.KillTokens(login); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetProfileByLogin(userLogin string, profileLogin string) (entity.UserProfile, error) {
	exists, err := s.repo.CheckLoginExists(profileLogin)
	if err != nil {
		return entity.UserProfile{}, err
	}

	if !exists {
		return entity.UserProfile{}, errors.ErrLoginDoesNotExist
	}

	userProfile, err := s.repo.GetProfile(profileLogin)
	if err != nil {
		return entity.UserProfile{}, err
	}

	if userLogin == profileLogin {
		return userProfile, nil
	}

	if userProfile.IsPublic {
		return userProfile, nil
	}

	friends, err := s.repo.ListFriends(profileLogin, 1000000, 0)
	if err != nil {
		return entity.UserProfile{}, err
	}

	if isFriend(friends, userLogin) {
		return userProfile, err
	}

	return entity.UserProfile{}, errors.ErrAccessDenied
}
