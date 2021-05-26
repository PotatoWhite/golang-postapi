package main

import (
	"log"
	"net/http"
	"os"
	app "postapi/app"
	"postapi/app/database"
)

func main() {
	// create app
	app := app.New()


	// create db connection
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)
	defer app.DB.Close()

	log.Println("Start Application")
	err = http.ListenAndServe(":8080", app.Router)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
