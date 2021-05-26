package app

import (
	"github.com/gorilla/mux"
	"postapi/app/database"
)

type App struct {
	Router *mux.Router
	DB 		database.PostDB
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/", app.IndexHandler()).Methods("GET")
	app.Router.HandleFunc("/posts", app.CreatePostHandler()).Methods("POST")
}

func New() *App{
	app := &App{
		Router: mux.NewRouter(),
	}

	app.initRoutes()
	return app
}