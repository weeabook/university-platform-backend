package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/weeabook/src/university-platform-backend/internal/config"
	"github.com/weeabook/src/university-platform-backend/internal/handler"
	"github.com/weeabook/src/university-platform-backend/internal/repository"
	"github.com/weeabook/src/university-platform-backend/internal/server"
	"github.com/weeabook/src/university-platform-backend/internal/service"
	"github.com/weeabook/src/university-platform-backend/pkg/database"
)

type App struct {
	cfg     *config.Config
	db      *sqlx.DB
	handler *handler.Handler
	server  *server.Server
}

func NewApp(cfg *config.Config) (*App, error) {
	db, err := database.NewPostgresDB(database.Config{
		Host:     cfg.DatabaseConfig.Host,
		Port:     cfg.DatabaseConfig.Port,
		Username: cfg.DatabaseConfig.Username,
		DBName:   cfg.DatabaseConfig.DBName,
		SSLMode:  cfg.DatabaseConfig.SSLMode,
		Password: cfg.DatabaseConfig.Password,
	})
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := server.NewServer(&http.Server{})

	return &App{
		cfg:     cfg,
		db:      db,
		handler: handler,
		server:  server,
	}, nil
}

func (a *App) Run() {
	go func() {
		if err := a.server.Run(a.cfg.Listen.Port, a.handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp shutdown")
	if err := a.server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutdown: %s", err.Error())
	}
	if err := a.db.Close(); err != nil {
		logrus.Errorf("error occured on close db-connection: %s", err.Error())
	}
}
