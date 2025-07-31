package service

import (
	"context"
	"time"

	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

type FriendService struct {
	userRepo       repository.User
	friendRepo     repository.Friend
	passwordHasher hasher.PasswordHasher
	signKey        string
	tokenTTL       time.Duration
}

func NewFriendService(userRepo repository.User, friendRepo repository.Friend, passwordHasher hasher.PasswordHasher, signKey string, tokenTTL time.Duration) *FriendService {
	return &FriendService{
		userRepo:       userRepo,
		friendRepo:     friendRepo,
		passwordHasher: passwordHasher,
		signKey:        signKey,
		tokenTTL:       tokenTTL,
	}
}

func (s *FriendService) AddFriend(ctx context.Context, userLogin, friendLogin string) error {
	isFriend, err := s.friendRepo.IsFriend(ctx, userLogin, friendLogin)
	if err != nil {
		return err
	}

	if isFriend {
		return nil
	}

	user, err := s.userRepo.GetUserByLogin(ctx, userLogin)
	if err != nil {
		if user == nil {
			return ErrUserNotFound
		}

		return err
	}

	return s.friendRepo.AddFriend(ctx, userLogin, friendLogin)
}

func (s *FriendService) RemoveFriend(ctx context.Context, userLogin, friendLogin string) error {
	isFriend, err := s.friendRepo.IsFriend(ctx, userLogin, friendLogin)
	if err != nil {
		return err
	}

	if !isFriend {
		return nil
	}

	user, err := s.userRepo.GetUserByLogin(ctx, userLogin)
	if err != nil {
		if user == nil {
			return ErrUserNotFound
		}

		return err
	}

	return s.friendRepo.RemoveFriend(ctx, userLogin, friendLogin)
}

func (s *FriendService) ListFriends(ctx context.Context, userLogin string, limit, offset int) ([]*FriendOutput, error) {
	friends, err := s.friendRepo.GetFriends(ctx, userLogin, limit, offset)
	if err != nil {
		return nil, err
	}

	friendsOutput := make([]*FriendOutput, 0, len(friends))
	for _, friend := range friends {
		friendOutput := &FriendOutput{
			UserLogin:   friend.UserLogin,
			FriendLogin: friend.FriendLogin,
			AddedAt:     friend.AddedAt,
		}

		friendsOutput = append(friendsOutput, friendOutput)
	}

	return friendsOutput, nil
}
