package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id            uuid.UUID `db:"id"`
	Content       string    `db:"content"`
	Author        string    `db:"author"`
	Tags          []string  `db:"tags"`
	CreatedAt     time.Time `db:"created_at"`
	LikesCount    int32     `db:"likes_count"`
	DislikesCount int32     `db:"dislikes_count"`
}
