package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cerande/go-reservation-app/src/pkg/config"
	"github.com/cerande/go-reservation-app/src/pkg/handlers"
	"github.com/cerande/go-reservation-app/src/pkg/renders"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	appConfig.UseCache = true
	appConfig.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsProduction

	appConfig.Session = session

	tc, err := renders.BuildCache()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.TemplateCache = tc
	renders.NewTemplate(&appConfig)
	handler := handlers.NewHandler(&appConfig)
	handlers.NewRepository(handler)

	server := &http.Server{
		Addr:    portNumber,
		Handler: registerRoutes(&appConfig),
	}

	fmt.Printf("Starting application on port %s \n", portNumber)
	err = server.ListenAndServe()
	log.Fatal(err)
}
