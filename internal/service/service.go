package service

import (
	"context"
	"time"

	"github.com/realdanielursul/pulse-api/internal/repo"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

// AuthService отвечает за аутентификацию и авторизацию
type Auth interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.UserProfile, error)
	SignIn(ctx context.Context, login, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (*models.UserClaims, error)
	UpdatePassword(ctx context.Context, userID string, oldPassword, newPassword string) error
}

// UserService отвечает за операции с пользователями
type User interface {
	GetProfile(ctx context.Context, login string, requesterID string) (*models.UserProfile, error)
	GetMyProfile(ctx context.Context, userID string) (*models.UserProfile, error)
	UpdateProfile(ctx context.Context, userID string, req *models.UpdateProfileRequest) (*models.UserProfile, error)
}

// CountryService отвечает за операции со странами
type Country interface {
	ListCountries(ctx context.Context, regions []string) ([]*models.Country, error)
	GetCountry(ctx context.Context, alpha2 string) (*models.Country, error)
}

// FriendsService отвечает за операции с друзьями
type Friend interface {
	AddFriend(ctx context.Context, userID, friendLogin string) error
	RemoveFriend(ctx context.Context, userID, friendLogin string) error
	ListFriends(ctx context.Context, userID string, limit, offset int) ([]*models.Friend, error)
}

// PostsService отвечает за операции с постами
type Post interface {
	CreatePost(ctx context.Context, userID string, req *models.CreatePostRequest) (*models.Post, error)
	GetPost(ctx context.Context, postID, requesterID string) (*models.Post, error)
	GetMyFeed(ctx context.Context, userID string, limit, offset int) ([]*models.Post, error)
	GetUserFeed(ctx context.Context, login, requesterID string, limit, offset int) ([]*models.Post, error)
	LikePost(ctx context.Context, postID, userID string) (*models.Post, error)
	DislikePost(ctx context.Context, postID, userID string) (*models.Post, error)
}

type ServicesDependencies struct {
	Repos  *repo.Repositories
	Hasher hasher.PasswordHasher

	SignKey  string
	TokenTTL time.Duration
}

type Services struct {
	Auth    Auth
	User    User
	Country Country
	Friend  Friend
	Post    Post
}

// type Service struct {
// 	repo           Repository
// 	passwordHasher hasher.SHA1Hasher
// 	signKey        string
// 	tokenTTL       time.Duration
// }
