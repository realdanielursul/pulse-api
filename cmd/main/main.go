package main

import (
	_ "github.com/lib/pq"
	"github.com/realdanielursul/pulse-api/config"
	"github.com/realdanielursul/pulse-api/pkg/logger"
	"github.com/realdanielursul/pulse-api/pkg/postgres"
	"github.com/sirupsen/logrus"
)

// SORT COUNTRIES

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

	////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////

	{
		// repo := repository.NewCountryRepository(db)

		// ctx := context.Background()

		// fmt.Println(repo.GetCountryByAlpha2(ctx, ""))
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////

	{
		// repo := repository.NewTokenRepository(db)

		// ctx := context.Background()

		// fmt.Println(repo.CreateToken(ctx, &entity.Token{"danixx", "token1", true}))
		// fmt.Println(repo.GetToken(ctx, "token1"))
		// fmt.Println(repo.InvalidateUserTokens(ctx, "danixx"))
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////////////////////////

	{
		// repo := repository.NewUserRepository(db)

		// ctx := context.Background()

		// login := "danixx"
		// email := "ursuldm@gmail.com"
		// passwordHash := "PASSWORD"
		// countryCode := "RU"
		// isPublic := true
		// phone := "+79219691565"
		// image := "https://link/to/image"

		// // CREATE
		// fmt.Println(repo.CreateUser(ctx, &entity.User{
		// 	Login:        login,
		// 	Email:        email,
		// 	PasswordHash: passwordHash,
		// 	CountryCode:  countryCode,
		// 	IsPublic:     isPublic,
		// 	Phone:        phone,
		// 	Image:        image,
		// }))

		// // GET BY LOGIN
		// fmt.Println(repo.GetUserByLogin(ctx, login))

		// // GET BY EMAIL
		// fmt.Println(repo.GetUserByEmail(ctx, email))

		// // GET BY PHONE
		// fmt.Println(repo.GetUserByPhone(ctx, phone))

		// // GET BY LOGIN AND PASSWORD
		// fmt.Println(repo.GetUserByLoginAndPassword(ctx, login, passwordHash))

		// // UPDATE
		// fmt.Println(repo.UpdateUser(ctx, login, nil, nil, nil, nil))

		// // UPDATE PASSWORD
		// fmt.Println(repo.UpdatePassword(ctx, login, "newPasswordHash"))

		// // GET BY LOGIN
		// fmt.Println(repo.GetUserByLogin(ctx, login))
	}
}
