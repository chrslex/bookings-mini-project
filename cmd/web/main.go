package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/chrslex/bookings-mini-project/pkg/config"
	"github.com/chrslex/bookings-mini-project/pkg/handlers"
	"github.com/chrslex/bookings-mini-project/pkg/render"
)

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// Set session in wide app config
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// App configuration
	app.TemplateCache = tc
	app.UseCache = false

	// Handle repository for handlers files
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Handle repository for render files
	render.NewTemplates(&app)

	// Start the service
	fmt.Printf(fmt.Sprintf("Starting application on port %s\n", portNumber))
	// http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
