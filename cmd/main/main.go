package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/realdanielursul/pulse-api/config"
	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/internal/service"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
	"github.com/realdanielursul/pulse-api/pkg/logger"
	"github.com/realdanielursul/pulse-api/pkg/postgres"
	"github.com/sirupsen/logrus"
)

// SORT COUNTRIES
// GEt POST BY ID (string or uuid)

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

	{
		authService := service.NewAuthService(deps.Repos.User, deps.Repos.Token, deps.Hasher, deps.SignKey, deps.TokenTTL)
		ctx := context.Background()

		fmt.Println(authService.Register(ctx, &service.AuthRegisterInput{
			Login:       "danixx",
			Email:       "ursuldm@gmail.com",
			Password:    "pizdaaaaa",
			CountryCode: "DE",
			IsPublic:    true,
			Phone:       "+79219691565",
			Image:       "https://link/to/image",
		}))
	}
}
