package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

const operationTimeout = 3 * time.Second

type User interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	GetUserByLoginAndPassword(ctx context.Context, login, passwordHash string) (*entity.User, error)
	UpdateUser(ctx context.Context, login string, countryCode, phone, image *string, isPublic *bool) error
	UpdatePassword(ctx context.Context, login, newPasswordHash string) error
}

type Token interface {
	CreateToken(ctx context.Context, token *entity.Token) error
	GetToken(ctx context.Context, tokenString string) (*entity.Token, error)
	InvalidateUserTokens(ctx context.Context, login string) error
}

type Country interface {
	GetAllCountries(ctx context.Context) ([]*entity.Country, error)
	GetCountriesByRegion(ctx context.Context, regions []string) ([]*entity.Country, error)
	GetCountryByAlpha2(ctx context.Context, alpha2 string) (*entity.Country, error)
}

type Friend interface {
	AddFriend(ctx context.Context, userLogin, friendLogin string) error
	RemoveFriend(ctx context.Context, userLogin, friendLogin string) error
	GetFriends(ctx context.Context, userLogin string, limit, offset int) ([]*entity.Friend, error)
	IsFriend(ctx context.Context, userLogin, friendLogin string) (bool, error)
}

type Post interface {
	CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
	GetPostById(ctx context.Context, postId string) (*entity.Post, error)
	GetUserPosts(ctx context.Context, userLogin string, limit, offset int) ([]*entity.Post, error)
	LikePost(ctx context.Context, postId, userLogin string) error
	DislikePost(ctx context.Context, postId, userLogin string) error
	GetPostReactionsCount(ctx context.Context, postId string) (likes, dislikes int, err error)
}

type Repositories struct {
	User
	Token
	Country
	Friend
	Post
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		User:    NewUserRepository(db),
		Token:   NewTokenRepository(db),
		Country: NewCountryRepository(db),
		Friend:  NewFriendRepository(db),
		Post:    NewPostRepository(db),
	}
}
