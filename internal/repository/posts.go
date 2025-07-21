package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) NewPost(login string, req models.NewPostRequest) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var post models.Post
	var likesCount sql.NullInt32
	var dislikesCount sql.NullInt32

	query := `INSERT INTO posts (id, content, author, tags) VALUES ($1, $2, $3, $4) RETURNING id, content, author, tags, created_at, array_length(likes_count, 1), array_length(dislikes_count, 1)`
	if err := r.db.QueryRowContext(ctx, query, uuid.New(), req.Content, login, pq.Array(req.Tags)).Scan(&post.Id, &post.Content, &post.Author, pq.Array(&post.Tags), &post.CreatedAt, &likesCount, &dislikesCount); err != nil {
		return models.Post{}, err
	}

	post.LikesCount = likesCount.Int32
	post.DislikesCount = dislikesCount.Int32

	return post, nil
}

func (r *PostgresRepository) GetPost(postId uuid.UUID) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var post models.Post
	var likesCount sql.NullInt32
	var dislikesCount sql.NullInt32

	query := `SELECT id, content, author, tags, created_at, array_length(likes_count, 1), array_length(dislikes_count, 1) FROM posts WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, postId).Scan(&post.Id, &post.Content, &post.Author, pq.Array(&post.Tags), &post.CreatedAt, &likesCount, &dislikesCount); err != nil {
		return models.Post{}, err
	}

	post.LikesCount = likesCount.Int32
	post.DislikesCount = dislikesCount.Int32

	return post, nil
}

func (r *PostgresRepository) ListPosts(login string, limit int, offset int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var posts []models.Post

	query := `SELECT id, content, author, tags, created_at, array_length(likes_count, 1), array_length(dislikes_count, 1) FROM posts WHERE author = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, login, limit, offset)
	if err != nil {
		return []models.Post{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var likesCount sql.NullInt32
		var dislikesCount sql.NullInt32

		if err := rows.Scan(&post.Id, &post.Content, &post.Author, pq.Array(&post.Tags), &post.CreatedAt, &likesCount, &dislikesCount); err != nil {
			return []models.Post{}, err
		}

		post.LikesCount = likesCount.Int32
		post.DislikesCount = dislikesCount.Int32

		posts = append(posts, post)
	}

	if rows.Err() != nil {
		return []models.Post{}, err
	}

	return posts, nil
}

func (r *PostgresRepository) LikePost(login string, postId uuid.UUID) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE posts SET likes_count = array_remove(likes_count, $1) WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, login, postId)
	if err != nil {
		return models.Post{}, err
	}

	query = `UPDATE posts SET dislikes_count = array_remove(dislikes_count, $1) WHERE id = $2`
	_, err = r.db.ExecContext(ctx, query, login, postId)
	if err != nil {
		return models.Post{}, err
	}

	query = `UPDATE posts SET likes_count = array_append(likes_count, $1) WHERE id = $2`
	_, err = r.db.ExecContext(ctx, query, login, postId)
	if err != nil {
		return models.Post{}, err
	}

	post, err := r.GetPost(postId)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r *PostgresRepository) DislikePost(login string, postId uuid.UUID) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE posts SET likes_count = array_remove(likes_count, $1) WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, login, postId)
	if err != nil {
		return models.Post{}, err
	}

	query = `UPDATE posts SET dislikes_count = array_remove(dislikes_count, $1) WHERE id = $2`
	_, err = r.db.ExecContext(ctx, query, login, postId)
	if err != nil {
		return models.Post{}, err
	}

	query = `UPDATE posts SET dislikes_count = array_append(dislikes_count, $1) WHERE id = $2`
	_, err = r.db.ExecContext(ctx, query, login, postId)
	if err != nil {
		return models.Post{}, err
	}

	post, err := r.GetPost(postId)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
