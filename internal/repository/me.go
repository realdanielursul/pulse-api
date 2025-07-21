package repository

import (
	"context"

	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (r *PostgresRepository) GetProfile(login string) (models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var userProfile models.UserProfile
	query := `SELECT login, email, country_code, is_public, phone, image FROM users WHERE login = $1`
	if err := r.db.QueryRowContext(ctx, query, login).Scan(&userProfile.Login, &userProfile.Email, &userProfile.CountryCode, &userProfile.IsPublic, &userProfile.Phone, &userProfile.Image); err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (r *PostgresRepository) UpdateProfile(login string, req models.UpdateProfileRequest) (models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	if req.CountryCode != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET country_code = $1 WHERE login = $2", *req.CountryCode, login)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.IsPublic != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET is_public = $1 WHERE login = $2", *req.IsPublic, login)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.Phone != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET phone = $1 WHERE login = $2", *req.Phone, login)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	if req.Image != nil {
		_, err := r.db.ExecContext(ctx, "UPDATE users SET image = $1 WHERE login = $2", *req.Image, login)
		if err != nil {
			return models.UserProfile{}, err
		}
	}

	userProfile, err := r.GetProfile(login)
	if err != nil {
		return models.UserProfile{}, err
	}

	return userProfile, nil
}

func (r *PostgresRepository) UpdatePassword(login string, req models.UpdatePasswordRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var oldPassword string

	query := `SELECT hash_password FROM users WHERE login = $1`
	if err := r.db.QueryRowContext(ctx, query, login).Scan(&oldPassword); err != nil {
		return err
	}

	if oldPassword != req.OldPassword {
		return errors.ErrInvalidOldPassword
	}

	query = `UPDATE users SET hash_password = $1 WHERE login = $2`
	_, err := r.db.ExecContext(ctx, query, req.NewPassword, login)

	return err
}
