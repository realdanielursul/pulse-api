package repository

import "github.com/jmoiron/sqlx"

type FriendRepository struct {
	*sqlx.DB
}

func NewFriendRepository(db *sqlx.DB) *FriendRepository {
	return &FriendRepository{db}
}
