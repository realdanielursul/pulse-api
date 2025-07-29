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

// SORT COUNTRIES
// GET POST BY ID (string or uuid)

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
		// userService := service.NewUserService(deps.Repos.User, deps.Repos.Friend, deps.Hasher, deps.SignKey, deps.TokenTTL)
		// ctx := context.Background()

		// fmt.Println(userService.UpdateProfile(ctx, "danixx", &service.UserUpdateProfileInput{
		// 	CountryCode: nil,
		// 	IsPublic:    nil,
		// 	Phone:       nil,
		// 	Image:       nil,
		// }))

		// fmt.Println(userService.GetMyProfile(ctx, "danixx"))

		// fmt.Println(userService.GetProfile(ctx, "danixx2", "danixx"))
	}

	{
		// authService := service.NewAuthService(deps.Repos.User, deps.Repos.Token, deps.Hasher, deps.SignKey, deps.TokenTTL)
		// ctx := context.Background()

		// fmt.Println(authService.Register(ctx, &service.AuthRegisterInput{
		// 	Login:       "danixx2",
		// 	Email:       "ursuldm@gmail.com2",
		// 	Password:    "pizdaaaaa",
		// 	CountryCode: "DE",
		// 	IsPublic:    false,
		// 	Phone:       "+792196915652",
		// 	Image:       "https://link/to/image",
		// }))

		// fmt.Println(authService.SignIn(ctx, &service.AuthSignInInput{
		// 	Login:    "danixx",
		// 	Password: "pizdaaaaa",
		// }))

		// fmt.Println(authService.ValidateToken(ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM4MDg2NDUsImlhdCI6MTc1MzgwMTQ0NSwiTG9naW4iOiJkYW5peHgifQ.5z09po2ovYqZ8TiqkA21KbtoL6LLoHan8vXFivWXa4c"))

		// fmt.Println(authService.UpdatePassword(ctx, "danixx", &service.AuthUpdatePasswordInput{
		// 	OldPassword: "pizdaaaaa",
		// 	NewPassword: "pizdaa",
		// }))
	}
}
