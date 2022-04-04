package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/siwodevilheart/bookings/pkg/config"
	"github.com/siwodevilheart/bookings/pkg/handlers"
	"github.com/siwodevilheart/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TmplCache = tc
	app.UseCache = false
	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	sv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = sv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
