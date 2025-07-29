package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	newUser := &entity.User{}
	sql := `INSERT INTO users (login, email, password_hash, country_code, is_public, phone, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING login, email, country_code, is_public, phone, image`
	if err := r.QueryRowContext(ctx, sql, user.Login, user.Email, user.PasswordHash, user.CountryCode, user.IsPublic, user.Phone, user.Image).Scan(&newUser.Login, &newUser.Email, &newUser.CountryCode, &newUser.IsPublic, &newUser.Phone, &newUser.Image); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *UserRepository) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	user := &entity.User{}
	sql := `SELECT login, email, country_code, is_public, phone, image FROM users WHERE login = $1`
	if err := r.QueryRowContext(ctx, sql, login).Scan(&user.Login, &user.Email, &user.CountryCode, &user.IsPublic, &user.Phone, &user.Image); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	user := &entity.User{}
	sql := `SELECT login, email, country_code, is_public, phone, image FROM users WHERE email = $1`
	if err := r.QueryRowContext(ctx, sql, email).Scan(&user.Login, &user.Email, &user.CountryCode, &user.IsPublic, &user.Phone, &user.Image); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	user := &entity.User{}
	sql := `SELECT login, email, country_code, is_public, phone, image FROM users WHERE phone = $1`
	if err := r.QueryRowContext(ctx, sql, phone).Scan(&user.Login, &user.Email, &user.CountryCode, &user.IsPublic, &user.Phone, &user.Image); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByLoginAndPassword(ctx context.Context, login, passwordHash string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	user := &entity.User{}
	sql := `SELECT login, email, country_code, is_public, phone, image FROM users WHERE login = $1 AND password_hash = $2`
	if err := r.QueryRowContext(ctx, sql, login, passwordHash).Scan(&user.Login, &user.Email, &user.CountryCode, &user.IsPublic, &user.Phone, &user.Image); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, login string, countryCode, phone, image *string, isPublic *bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	if countryCode != nil {
		_, err := r.ExecContext(ctx, "UPDATE users SET country_code = $1 WHERE login = $2", *countryCode, login)
		if err != nil {
			return err
		}
	}

	if isPublic != nil {
		_, err := r.ExecContext(ctx, "UPDATE users SET is_public = $1 WHERE login = $2", *isPublic, login)
		if err != nil {
			return err
		}
	}

	if phone != nil {
		_, err := r.ExecContext(ctx, "UPDATE users SET phone = $1 WHERE login = $2", *phone, login)
		if err != nil {
			return err
		}
	}

	if image != nil {
		_, err := r.ExecContext(ctx, "UPDATE users SET image = $1 WHERE login = $2", *image, login)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, login, newPasswordHash string) error {
	sql := `UPDATE users SET password_hash = $1 WHERE login = $2`
	_, err := r.ExecContext(ctx, sql, newPasswordHash, login)

	return err
}
