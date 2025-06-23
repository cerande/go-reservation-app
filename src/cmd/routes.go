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

	fileServer := http.FileServer(http.Dir("./src/static/"))
	mux.Handle("/src/static/*", http.StripPrefix("/src/static", fileServer))

	return mux
}
