package entity

import "time"

type Friend struct {
	FriendLogin string    `db:"friend_login"`
	AddedAt     time.Time `db:"added_at"`
}
