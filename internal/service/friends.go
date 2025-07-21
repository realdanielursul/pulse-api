package service

import (
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) AddFriend(userLogin string, login string) error {
	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return err
	}

	if !exists {
		return errors.ErrLoginDoesNotExist
	}

	return s.repo.AddFriend(userLogin, login)
}

func (s *Service) RemoveFriend(userLogin string, login string) error {
	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return err
	}

	if !exists {
		return errors.ErrLoginDoesNotExist
	}

	return s.repo.RemoveFriend(userLogin, login)
}

func (s *Service) ListFriends(login string, limit int, offset int) ([]models.FriendInfo, error) {
	if limit < 0 || offset < 0 {
		return []models.FriendInfo{}, errors.ErrInvalidPaginationParams
	}

	return s.repo.ListFriends(login, limit, offset)
}
