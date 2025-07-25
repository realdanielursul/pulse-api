package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/ursulgwopp/pulse-api/config"
	"github.com/ursulgwopp/pulse-api/internal/handler"
	"github.com/ursulgwopp/pulse-api/internal/repository"
	"github.com/ursulgwopp/pulse-api/internal/service"
	"github.com/ursulgwopp/pulse-api/pkg/httpserver"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.NewConfig("./config/local.yaml")
	if err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config.Postgres{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.Database,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewPostgresRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

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
