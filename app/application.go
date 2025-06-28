package app

import (
	"learning/app/controllers"
)

type Application struct {
	Server *controllers.Server
}

func New() *Application {
	server := controllers.New()
	return &Application{
		Server: server,
	}
}

func (app *Application) Run() {
   app.Server.Run()
}
