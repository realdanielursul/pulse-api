package repository

import (
	"context"
	"time"

	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) AddFriend(userLogin string, login string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	if userLogin == login {
		return nil
	}

	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM friends WHERE login = $1 AND EXISTS (SELECT 1 FROM UNNEST(friends_info) AS fi WHERE fi.login = $2))`
	if err := r.db.QueryRowContext(ctx, query, userLogin, login).Scan(&exists); err != nil {
		return err
	}

	if exists {
		return nil
	}

	query = `UPDATE friends SET friends_info = ARRAY_APPEND(friends_info, ROW($1, $2::TIMESTAMP)::friend_info) WHERE login = $3`
	_, err := r.db.ExecContext(ctx, query, login, time.Now(), userLogin)

	return err
}

func (r *PostgresRepository) RemoveFriend(userLogin string, login string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE friends SET friends_info = (SELECT ARRAY_AGG(fi) FROM UNNEST(friends_info) AS fi WHERE fi.login <> $1) WHERE login = $2`
	_, err := r.db.ExecContext(ctx, query, login, userLogin)

	return err
}

func (r *PostgresRepository) ListFriends(login string, limit int, offset int) ([]models.FriendInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var friends []models.FriendInfo

	query := `WITH unnested_friends AS (SELECT (UNNEST(friends_info)).* AS friend_info FROM friends WHERE login = $1) SELECT * FROM unnested_friends LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, login, limit, offset)
	if err != nil {
		return []models.FriendInfo{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var friend models.FriendInfo
		if err := rows.Scan(&friend.Login, &friend.AddedAt); err != nil {
			return []models.FriendInfo{}, err
		}

		friends = append(friends, friend)
	}

	if rows.Err() != nil {
		return []models.FriendInfo{}, err
	}

	return friends, nil
}
