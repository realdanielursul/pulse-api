package service

import (
	"github.com/google/uuid"
	"github.com/ursulgwopp/pulse-api/internal/entity"
)

type Repository interface {
	ListCountries(regions []string) ([]entity.Country, error)
	GetCountryByAlpha2(alpha2 string) (entity.Country, error)

	Register(req entity.RegisterRequest) (entity.UserProfile, error)
	SignIn(req entity.SignInRequest) (string, error)
	AddToken(login string, token string) error
	ValidateToken(token string) error
	KillTokens(login string) error

	GetProfile(login string) (entity.UserProfile, error)
	UpdateProfile(login string, req entity.UpdateProfileRequest) (entity.UserProfile, error)
	UpdatePassword(login string, req entity.UpdatePasswordRequest) error

	AddFriend(userLogin string, login string) error
	RemoveFriend(userLogin string, login string) error
	ListFriends(login string, limit int, offset int) ([]entity.FriendInfo, error)

	NewPost(login string, req entity.NewPostRequest) (entity.Post, error)
	GetPost(postId uuid.UUID) (entity.Post, error)
	ListPosts(login string, limit int, offset int) ([]entity.Post, error)
	LikePost(login string, postId uuid.UUID) (entity.Post, error)
	DislikePost(login string, postId uuid.UUID) (entity.Post, error)

	CheckLoginExists(login string) (bool, error)
	CheckEmailExists(email string) (bool, error)
	CheckCountryCodeExists(alpha2 string) (bool, error)
	CheckPhoneExists(phone string) (bool, error)
	// CheckUserIdByLogin(login string) (int, error)
	// CheckLoginByUserId(id int) (string, error)
	CheckProfileIsPublic(login string) (bool, error)
	CheckPostIdExists(id uuid.UUID) (bool, error)
	CheckPostAuthor(id uuid.UUID) (string, error)
	// CheckUserIdExists(id int) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
