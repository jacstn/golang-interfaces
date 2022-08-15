package main

import (
	"github.com/go-chi/chi"
	"github.com/jacstn/golang-playground/simple-server/pkg/handlers"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}
