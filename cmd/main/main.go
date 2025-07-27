package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/realdanielursul/pulse-api/config"
	v1 "github.com/realdanielursul/pulse-api/internal/controller/http/v1"
	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/internal/service"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
	"github.com/realdanielursul/pulse-api/pkg/httpserver"
	"github.com/realdanielursul/pulse-api/pkg/logger"
	"github.com/realdanielursul/pulse-api/pkg/postgres"
	"github.com/sirupsen/logrus"
)

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

	repository := repository.NewPostgresRepository(db)
	service := service.NewService(repository, hasher.NewSHA1Hasher(cfg.Hasher.Salt), cfg.JWT.SignKey, cfg.JWT.TokenTTL)
	handler := v1.NewHandler(service)

	srv := &httpserver.Server{}
	go func() {
		if err := srv.Run(cfg.HTTP.Port, handler.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()

	logrus.Printf("App '%s %s' Started", cfg.App.Name, cfg.App.Version)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	logrus.Printf("App '%s %s' Shutted Down", cfg.App.Name, cfg.App.Version)
}
