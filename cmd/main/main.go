package main

import (
	_ "github.com/lib/pq"
	"github.com/realdanielursul/pulse-api/config"
	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/internal/service"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
	"github.com/realdanielursul/pulse-api/pkg/logger"
	"github.com/realdanielursul/pulse-api/pkg/postgres"
	"github.com/sirupsen/logrus"
)

// GET POST BY ID (string or uuid?)
// RENAME FRIENDS FIELDS (maybe follower?)

func main() {
	logger.SetLogrus()

	cfg, err := config.NewConfig("./config/local.yaml")
	if err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repositories := repository.NewRepositories(db)

	deps := service.ServicesDependencies{
		Repos:    repositories,
		Hasher:   hasher.NewSHA1Hasher(cfg.Hasher.Salt),
		SignKey:  cfg.JWT.SignKey,
		TokenTTL: cfg.JWT.TokenTTL,
	}

	service := service.NewService(deps)

}
