package main

import (
	"net/http"

	"github.com/Xcoder2088/Location/pkg/handlers"
	"github.com/Xcoder2088/Location/pkg/handlers/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteInConsole)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./Static/"))
	mux.Handle("/Static/*", http.StripPrefix("/Static", fileServer))
	return mux
}
