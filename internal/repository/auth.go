package repository

import (
	"context"

	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) Register(req models.RegisterRequest) (models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var userProfile models.UserProfile

	query := `INSERT INTO users (login, email, hash_password, country_code, is_public, phone, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING login, email, country_code, is_public, phone, image`
	if err := r.db.QueryRowContext(ctx, query, req.Login, req.Email, req.Password, req.CountryCode, req.IsPublic, req.Phone, req.Image).Scan(&userProfile.Login, &userProfile.Email, &userProfile.CountryCode, &userProfile.IsPublic, &userProfile.Phone, &userProfile.Image); err != nil {
		return models.UserProfile{}, err
	}

	query = `INSERT INTO friends (login) VALUES ($1)`
	_, err := r.db.ExecContext(ctx, query, userProfile.Login)
	if err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (r *PostgresRepository) SignIn(req models.SignInRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE login = $1 AND hash_password = $2)`
	if err := r.db.QueryRowContext(ctx, query, req.Login, req.Password).Scan(&exists); err != nil {
		return "", err
	}

	if !exists {
		return "", errors.ErrInvalidUsernameOrPassword
	}

	return req.Login, nil
}

func (r *PostgresRepository) AddToken(login string, token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `INSERT INTO tokens (login, token, is_valid) VALUES ($1, $2, TRUE)`
	_, err := r.db.ExecContext(ctx, query, login, token)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) ValidateToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var is_valid bool

	query := `SELECT is_valid FROM tokens WHERE token = $1`
	if err := r.db.QueryRowContext(ctx, query, token).Scan(&is_valid); err != nil {
		return err
	}

	if !is_valid {
		return errors.ErrInvalidToken
	}

	return nil
}

func (r *PostgresRepository) KillTokens(login string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE tokens SET is_valid = FALSE WHERE login = $1`
	_, err := r.db.ExecContext(ctx, query, login)

	return err
}
