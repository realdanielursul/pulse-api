package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

type FriendRepository struct {
	*sqlx.DB
}

func NewFriendRepository(db *sqlx.DB) *FriendRepository {
	return &FriendRepository{db}
}

func (r *FriendRepository) AddFriend(ctx context.Context, userLogin, friendLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	sql := `INSERT INTO friends (user_login, friend_login) VALUES ($1, $2)`
	_, err := r.ExecContext(ctx, sql, userLogin, friendLogin)

	return err
}

func (r *FriendRepository) RemoveFriend(ctx context.Context, userLogin, friendLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	sql := `DELETE FROM friends WHERE user_login = $1 AND friend_login = $2`
	_, err := r.ExecContext(ctx, sql, userLogin, friendLogin)

	return err
}

func (r *FriendRepository) GetFriends(ctx context.Context, userLogin string, limit, offset int) ([]*entity.Friend, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	friends := make([]*entity.Friend, 0, 100)
	sql := `SELECT * FROM friends WHERE user_login = $1 ORDER BY added_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.QueryContext(ctx, sql, userLogin, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		friend := &entity.Friend{}
		if rows.Scan(&friend.UserLogin, &friend.FriendLogin, &friend.AddedAt); err != nil {
			return nil, err
		}

		friends = append(friends, friend)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return friends, nil
}

func (r *FriendRepository) IsFriend(ctx context.Context, userLogin, friendLogin string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM friends WHERE user_login = $1 AND friend_login = $2)`
	if err := r.QueryRowContext(ctx, query, userLogin, friendLogin).Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return exists, nil
}
