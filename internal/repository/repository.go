package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

var operationTimeout = time.Second * 3

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
