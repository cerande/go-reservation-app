package main

import (
	"net/http"

	"github.com/cerande/go-reservation-app/src/pkg/config"
	"github.com/cerande/go-reservation-app/src/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func registerRoutes(cfg *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)
	mux.Use(sessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/make-reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./src/static/"))
	mux.Handle("/src/static/*", http.StripPrefix("/src/static", fileServer))

	return mux
}
