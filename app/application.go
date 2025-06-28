package app

import (
	"learning/app/controllers"
	"learning/config"

	"github.com/gin-gonic/gin"
	"github.com/roonglit/credentials/pkg/credentials"
)

type Application struct {
	Server *controllers.Server
}

func New() *Application {
	config := loadConfig()

	server := controllers.New(
config)
	return &Application{
		Server: server,
	}
}

func (app *Application) Run() {
	app.Server.Run()
}

func loadConfig() *config.Config {
	reader := credentials.NewConfigReader()

	var config config.Config

	if err := reader.Read(gin.Mode(), &config); err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	return &config
}
