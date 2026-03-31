package app

import (
	"assign1/internal/config"
	"assign1/internal/database/repositories"
)

type App struct {
	Repos *repositories.Repositories
	Cfg   *config.Settings
}

func NewApp(repos *repositories.Repositories, cfg *config.Settings) *App {
	return &App{
		Repos: repos,
		Cfg:   cfg,
	}
}
