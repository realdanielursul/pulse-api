package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

type TokenRepository struct {
	*sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db}
}

func (r *TokenRepository) CreateToken(ctx context.Context, token *entity.Token) error {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	sql := `INSERT INTO tokens (login, token_string, is_valid) VALUES ($1, $2, $3)`
	_, err := r.ExecContext(ctx, sql, token.Login, token.TokenString, token.IsValid)

	return err
}

func (r *TokenRepository) GetToken(ctx context.Context, tokenString string) (*entity.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	token := &entity.Token{}
	sql := `SELECT * FROM tokens WHERE token_string = $1`
	if err := r.QueryRowContext(ctx, sql, tokenString).Scan(&token.Login, &token.TokenString, &token.IsValid); err != nil {
		return nil, err
	}

	return token, nil
}

func (r *TokenRepository) InvalidateUserTokens(ctx context.Context, login string) error {
	ctx, cancel := context.WithTimeout(ctx, operationTimeout)
	defer cancel()

	sql := `UPDATE tokens SET is_valid = FALSE WHERE login = $1`
	_, err := r.ExecContext(ctx, sql, login)

	return err
}
