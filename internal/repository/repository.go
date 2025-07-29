package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

const operationTimeout = 3 * time.Second

// UserRepository отвечает за хранение и получение данных о пользователях
type User interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	GetUserByLoginAndPassword(ctx context.Context, login, passwordHash string) (*entity.User, error)
	UpdateUser(ctx context.Context, login string, countryCode, phone, image *string, isPublic *bool) error
	UpdatePassword(ctx context.Context, login, newPasswordHash string) error
}

// TokenRepository отвечает за хранение и валидацию токенов
type Token interface {
	CreateToken(ctx context.Context, token *entity.Token) error
	GetToken(ctx context.Context, tokenString string) (*entity.Token, error)
	InvalidateUserTokens(ctx context.Context, login string) error
}

// CountryRepository отвечает за хранение и получение данных о странах
type Country interface {
	GetAllCountries(ctx context.Context) ([]*entity.Country, error)
	GetCountriesByRegion(ctx context.Context, regions []string) ([]*entity.Country, error)
	GetCountryByAlpha2(ctx context.Context, alpha2 string) (*entity.Country, error)
}

// FriendsRepository отвечает за хранение и получение данных о друзьях
type Friend interface {
	AddFriend(ctx context.Context, userID, friendID string) error
	RemoveFriend(ctx context.Context, userID, friendID string) error
	GetFriends(ctx context.Context, userID string, limit, offset int) ([]*entity.Friend, error)
	IsFriend(ctx context.Context, userID, friendID string) (bool, error)
}

// PostsRepository отвечает за хранение и получение данных о постах
type Post interface {
	CreatePost(ctx context.Context, post *entity.Post) error
	GetPostByID(ctx context.Context, postID string) (*entity.Post, error)
	GetUserPosts(ctx context.Context, userID string, limit, offset int) ([]*entity.Post, error)
	LikePost(ctx context.Context, postID, userID string) error
	DislikePost(ctx context.Context, postID, userID string) error
	GetPostReactionsCount(ctx context.Context, postID string) (likes, dislikes int, err error)
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
		User:  NewUserRepository(db),
		Token: NewTokenRepository(db),
		// Country: NewCountryRepository(db),
		// Friend:  NewFriendRepository(db),
		// Post:    NewPostRepository(db),
	}
}
