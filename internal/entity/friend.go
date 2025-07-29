package entity

import "time"

type Friend struct {
	UserLogin   string    `db:"user_login"`
	FriendLogin string    `db:"friend_login"`
	AddedAt     time.Time `db:"added_at"`
}
