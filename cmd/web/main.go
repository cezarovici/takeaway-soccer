package main

import (
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cezarovici/takeaway-soccer/packages/config"
	"github.com/cezarovici/takeaway-soccer/packages/handler"
	"github.com/cezarovici/takeaway-soccer/packages/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, errCreatingTemlateCache := render.CreateTemplateCache()
	if errCreatingTemlateCache != nil {
		os.Exit(1)
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplate(&app)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := server.ListenAndServe()
	if err != nil {
		os.Exit(2)
	}
}
