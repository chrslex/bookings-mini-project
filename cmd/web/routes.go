package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/chrslex/bookings-mini-project/pkg/config"
	"github.com/chrslex/bookings-mini-project/pkg/handlers"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(cfg *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
