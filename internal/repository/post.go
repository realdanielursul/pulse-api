package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

type PostRepository struct {
	*sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	newPost := &entity.Post{LikesCount: 0, DislikesCount: 0}
	sql := `INSERT INTO posts (content, author, tags) VALUES ($1, $2, $3) RETURNING id, content, author, tags, created_at`
	if err := r.QueryRowContext(ctx, sql, post.Content, post.Author, pq.Array(post.Tags)).Scan(&newPost.Id, &newPost.Content, &newPost.Author, pq.Array(&newPost.Tags), &newPost.CreatedAt); err != nil {
		return nil, err
	}

	return newPost, nil
}

func (r *PostRepository) GetPostById(ctx context.Context, postId string) (*entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	post := &entity.Post{}
	sql := `SELECT * FROM posts WHERE id = $1`
	if err := r.QueryRowContext(ctx, sql, postId).Scan(&post.Id, &post.Content, &post.Author, pq.Array(&post.Tags), &post.CreatedAt); err != nil {
		return nil, err
	}

	likes, dislikes, err := r.GetPostReactionsCount(ctx, postId)
	if err != nil {
		return nil, err
	}

	post.LikesCount = likes
	post.DislikesCount = dislikes

	return post, nil
}

func (r *PostRepository) GetUserPosts(ctx context.Context, userLogin string, limit, offset int) ([]*entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	posts := make([]*entity.Post, 0, 100)
	sql := `SELECT * FROM posts WHERE author = $1`
	rows, err := r.QueryContext(ctx, sql, userLogin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &entity.Post{}
		if err := rows.Scan(&post.Id, &post.Content, &post.Author, pq.Array(&post.Tags), &post.CreatedAt); err != nil {
			return nil, err
		}

		likes, dislikes, err := r.GetPostReactionsCount(ctx, post.Id.String())
		if err != nil {
			return nil, err
		}

		post.LikesCount = likes
		post.DislikesCount = dislikes

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) LikePost(ctx context.Context, postId, userLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	tx, err := r.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		DELETE FROM post_reactions 
		WHERE post_id = $1 AND user_login = $2 AND reaction_type = 'dislike'`,
		postId, userLogin)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO post_reactions (post_id, user_login, reaction_type, created_at)
		VALUES ($1, $2, 'like', $3)
		ON CONFLICT (post_id, user_login) DO UPDATE 
		SET reaction_type = 'like', created_at = $3`,
		postId, userLogin, time.Now())
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostRepository) DislikePost(ctx context.Context, postId, userLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	tx, err := r.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		DELETE FROM post_reactions 
		WHERE post_id = $1 AND user_login = $2 AND reaction_type = 'like'`,
		postId, userLogin)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO post_reactions (post_id, user_login, reaction_type, created_at)
		VALUES ($1, $2, 'dislike', $3)
		ON CONFLICT (post_id, user_login) DO UPDATE 
		SET reaction_type = 'dislike', created_at = $3`,
		postId, userLogin, time.Now())
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostRepository) GetPostReactionsCount(ctx context.Context, postId string) (likes, dislikes int, err error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	sql := `SELECT COUNT(CASE WHEN reaction_type = 'like' THEN 1 END) as likes, COUNT(CASE WHEN reaction_type = 'dislike' THEN 1 END) as dislikes FROM post_reactions WHERE post_id = $1`
	if err := r.QueryRowContext(ctx, sql, postId).Scan(&likes, &dislikes); err != nil {
		return -1, -1, err
	}

	return likes, dislikes, nil
}
