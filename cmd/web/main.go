package main

import (
	"bookings_app/pkg/config"
	"bookings_app/pkg/handlers"
	"bookings_app/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

func main() {

	//change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 5 // session lasts 5 hours
	// cookies will be default for session
	session.Cookie.Persist = true //after closed, session keeps persisting
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // no https, no encryption

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc

	app.TemplateCache = tc
	// We allow a dynamic rebuild of our templates
	// but can be set to false to not rebuild each
	// time.
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// here we pass the url and function
	//http.HandleFunc("/", handlers.Repo.Home) // here we pass Repo to receiver of Home func
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on %v\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)

	// Instead of calling http.HandleFunc directly as above
	// we use routes.go to server the pages via the routes
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
