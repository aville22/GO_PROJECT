package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/aville22/greeneats/internal/config"
	"github.com/aville22/greeneats/internal/handlers"
	"github.com/aville22/greeneats/internal/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{Addr: ":8080", Handler: routes(&app)}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
func run() error {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCash()
	if err != nil {
		log.Fatal(err)
		return err
	}
	app.TemplateCache = tc
	app.UseCash = false
	render.NewTemplate(&app) // Перенесено сюда
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//_ = http.ListenAndServe(":8080", nil)

	return nil
}
