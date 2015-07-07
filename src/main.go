package main

import (
	app "github.com/dblokhin/webapp"
	"interfaces/handlers"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Server is preparing to start")
	Application := app.GetApplication()

	if Application.Config.Site.Disabled {
		log.Println("Site is disabled")
		Application.Routes(app.MapRoutes{{"/": handlers.HandleDisabled{}}})
	} else {
		Application.Routes(app.MapRoutes{
			{"/": handlers.HandleHome{}},
			//"/else": handlers.HandleElse{},
			{"/{url:.*}": handlers.Handle404{}},
		})
	}

	Application.Run()
	log.Println("Exit")
}
