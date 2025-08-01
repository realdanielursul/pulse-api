package service

import (
	"context"

	"github.com/realdanielursul/pulse-api/internal/repository"
)

type FriendService struct {
	userRepo   repository.User
	friendRepo repository.Friend
}

func NewFriendService(userRepo repository.User, friendRepo repository.Friend) *FriendService {
	return &FriendService{
		userRepo:   userRepo,
		friendRepo: friendRepo,
	}
}

func (s *FriendService) AddFriend(ctx context.Context, userLogin, friendLogin string) error {
	user, err := s.userRepo.GetUserByLogin(ctx, userLogin)
	if err != nil {
		if user == nil {
			return ErrUserNotFound
		}

		return err
	}

	isFriend, err := s.friendRepo.IsFriend(ctx, userLogin, friendLogin)
	if err != nil {
		return err
	}

	if isFriend {
		return nil
	}

	return s.friendRepo.AddFriend(ctx, userLogin, friendLogin)
}

func (s *FriendService) RemoveFriend(ctx context.Context, userLogin, friendLogin string) error {
	user, err := s.userRepo.GetUserByLogin(ctx, userLogin)
	if err != nil {
		if user == nil {
			return ErrUserNotFound
		}

		return err
	}

	isFriend, err := s.friendRepo.IsFriend(ctx, userLogin, friendLogin)
	if err != nil {
		return err
	}

	if !isFriend {
		return nil
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
			FriendLogin: friend.FriendLogin,
			AddedAt:     friend.AddedAt,
		}

		friendsOutput = append(friendsOutput, friendOutput)
	}

	return friendsOutput, nil
}
