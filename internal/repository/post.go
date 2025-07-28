package repository

import "github.com/jmoiron/sqlx"

type PostRepository struct {
	*sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db}
}
