package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Country struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Region string `json:"region"`
}

type RegisterRequest struct {
	Login       string `json:"login" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CountryCode string `json:"countryCode" binding:"required"`
	IsPublic    bool   `json:"isPublic" binding:"required"`
	Phone       string `json:"phone,omitempty"`
	Image       string `json:"image,omitempty"`
}

type SignInRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Login string
}

type UserProfile struct {
	Login       string `json:"login" binding:"required"`
	Email       string `json:"email" binding:"required"`
	CountryCode string `json:"countryCode" binding:"required"`
	IsPublic    bool   `json:"isPublic" binding:"required"`
	Phone       string `json:"phone,omitempty"`
	Image       string `json:"image,omitempty"`
}

type UpdateProfileRequest struct {
	CountryCode *string `json:"countryCode"`
	IsPublic    *bool   `json:"isPublic"`
	Phone       *string `json:"phone"`
	Image       *string `json:"image"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type FriendInfo struct {
	Login   string    `json:"login"`
	AddedAt time.Time `json:"addedAt"`
}

type LoginRequest struct {
	Login string `json:"login" binding:"required"`
}

type NewPostRequest struct {
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags" binding:"required"`
}

type Post struct {
	Id            uuid.UUID `json:"id"`
	Content       string    `json:"content"`
	Author        string    `json:"author"`
	Tags          []string  `json:"tags"`
	CreatedAt     time.Time `json:"createdAt"`
	LikesCount    int32     `json:"likesCount"`
	DislikesCount int32     `json:"dislikesCount"`
}
