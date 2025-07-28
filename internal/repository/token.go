package repository

import "github.com/jmoiron/sqlx"

type TokenRepository struct {
	*sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db}
}
