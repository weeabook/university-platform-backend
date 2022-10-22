package app

import (
	"github.com/weeabook/src/university-platform-backend/internal/config"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) (*App, error) {
	return &App{
		cfg: cfg,
	}, nil
}

func (a *App) Run() {}
