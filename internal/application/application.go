package application

import "github.com/Shin-Thant/service-test/internal/config"

type Application struct {
	config *config.Config
}

func New(c *config.Config) *Application {
	return &Application{
		config: c,
	}
}
