package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type AuthRegisterInput struct {
	Login       string
	Email       string
	Password    string
	CountryCode string
	IsPublic    bool
	Phone       string
	Image       string
}

type AuthRegisterOutput struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	CountryCode string `json:"country_code"`
	IsPublic    bool   `json:"is_public"`
	Phone       string `json:"phone,omitempty"`
	Image       string `json:"image,omitempty"`
}

type AuthSignInInput struct {
	Login    string
	Password string
}

type AuthUpdatePasswordInput struct {
	OldPassword string
	NewPassword string
}

type Auth interface {
	Register(ctx context.Context, input *AuthRegisterInput) (*AuthRegisterOutput, error)
	SignIn(ctx context.Context, input *AuthSignInInput) (string, error)
	ValidateToken(ctx context.Context, tokenString string) (bool, error)
	UpdatePassword(ctx context.Context, userLogin string, input *AuthUpdatePasswordInput) error
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type UserOutput struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	CountryCode string `json:"country_code"`
	IsPublic    bool   `json:"is_public"`
	Phone       string `json:"phone,omitempty"`
	Image       string `json:"image,omitempty"`
}

type UserUpdateProfileInput struct {
	CountryCode *string
	IsPublic    *bool
	Phone       *string
	Image       *string
}

type User interface {
	GetProfile(ctx context.Context, login string, requesterLogin string) (*UserOutput, error)
	GetMyProfile(ctx context.Context, userLogin string) (*UserOutput, error)
	UpdateProfile(ctx context.Context, userLogin string, input *UserUpdateProfileInput) (*UserOutput, error)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type CountryOutput struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Region string `json:"region"`
}

type Country interface {
	ListCountries(ctx context.Context, regions []string) ([]*CountryOutput, error)
	GetCountry(ctx context.Context, alpha2 string) (*CountryOutput, error)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type FriendOutput struct {
	UserLogin   string    `json:"user_login"`
	FriendLogin string    `json:"friend_login"`
	AddedAt     time.Time `json:"added_at"`
}

type Friend interface {
	AddFriend(ctx context.Context, userLogin, friendLogin string) error
	RemoveFriend(ctx context.Context, userLogin, friendLogin string) error
	ListFriends(ctx context.Context, userLogin string, limit, offset int) ([]*Friend, error)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type PostCreateInput struct {
	Content string
	Author  string
	Tags    []string
}

type PostOutput struct {
	Id            uuid.UUID `json:"id"`
	Content       string    `json:"content"`
	Author        string    `json:"author"`
	Tags          []string  `json:"tags"`
	CreatedAt     time.Time `json:"created_at"`
	LikesCount    int       `json:"likes_count"`
	DislikesCount int       `json:"dislikes_count"`
}

type Post interface {
	CreatePost(ctx context.Context, req *PostCreateInput) (*PostOutput, error)
	GetPost(ctx context.Context, postID, requesterLogin string) (*PostOutput, error)
	GetMyFeed(ctx context.Context, userLogin string, limit, offset int) ([]*PostOutput, error)
	GetUserFeed(ctx context.Context, login, requesterLogin string, limit, offset int) ([]*PostOutput, error)
	LikePost(ctx context.Context, postID, userLogin string) (*PostOutput, error)
	DislikePost(ctx context.Context, postID, userLogin string) (*PostOutput, error)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ServicesDependencies struct {
	Repos  *repository.Repositories
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
