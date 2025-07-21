package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *PostgresRepository) CheckLoginExists(login string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)`
	if err := r.db.QueryRowContext(ctx, query, login).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckEmailExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	if err := r.db.QueryRowContext(ctx, query, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckCountryCodeExists(alpha2 string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM countries WHERE alpha2 = $1)`
	if err := r.db.QueryRowContext(ctx, query, alpha2).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckPhoneExists(phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE phone = $1)`
	if err := r.db.QueryRowContext(ctx, query, phone).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckUserIdByLogin(login string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var id int

	query := `SELECT id FROM users WHERE login = $1`
	if err := r.db.QueryRowContext(ctx, query, login).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) CheckLoginByUserId(id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var login string

	query := `SELECT login FROM users WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&login); err != nil {
		return "", err
	}

	return login, nil
}

func (r *PostgresRepository) CheckProfileIsPublic(login string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var is_public bool

	query := `SELECT is_public FROM users WHERE login = $1`
	if err := r.db.QueryRowContext(ctx, query, login).Scan(&is_public); err != nil {
		return false, err
	}

	return is_public, nil
}

func (r *PostgresRepository) CheckPostIdExists(id uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM posts WHERE id = $1)`
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *PostgresRepository) CheckPostAuthor(id uuid.UUID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var author string

	query := `SELECT author FROM posts WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, id).Scan(&author); err != nil {
		return "", err
	}

	return author, nil
}
