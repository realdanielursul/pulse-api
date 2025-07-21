package service

import (
	"github.com/google/uuid"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

type Repository interface {
	ListCountries(regions []string) ([]models.Country, error)
	GetCountryByAlpha2(alpha2 string) (models.Country, error)

	Register(req models.RegisterRequest) (models.UserProfile, error)
	SignIn(req models.SignInRequest) (string, error)
	AddToken(login string, token string) error
	ValidateToken(token string) error
	KillTokens(login string) error

	GetProfile(login string) (models.UserProfile, error)
	UpdateProfile(login string, req models.UpdateProfileRequest) (models.UserProfile, error)
	UpdatePassword(login string, req models.UpdatePasswordRequest) error

	AddFriend(userLogin string, login string) error
	RemoveFriend(userLogin string, login string) error
	ListFriends(login string, limit int, offset int) ([]models.FriendInfo, error)

	NewPost(login string, req models.NewPostRequest) (models.Post, error)
	GetPost(postId uuid.UUID) (models.Post, error)
	ListPosts(login string, limit int, offset int) ([]models.Post, error)
	LikePost(login string, postId uuid.UUID) (models.Post, error)
	DislikePost(login string, postId uuid.UUID) (models.Post, error)

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
