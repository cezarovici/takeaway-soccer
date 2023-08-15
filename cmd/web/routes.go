package main

import (
	"net/http"

	"github.com/cezarovici/takeaway-soccer/packages/config"
	"github.com/cezarovici/takeaway-soccer/packages/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(SessionLoad)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)

	mux.Get("/", http.HandlerFunc(handler.Repo.HandleHomePage))
	mux.Get("/adauga-editie", http.HandlerFunc(handler.Repo.HandleAddEditie))

	apiRouter := chi.NewRouter()
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(middleware.NoCache)
	apiRouter.Post("/adauga-editie-json", http.HandlerFunc(handler.Repo.HandleAdaugaEditieJson))
	mux.Mount("/api", apiRouter)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
