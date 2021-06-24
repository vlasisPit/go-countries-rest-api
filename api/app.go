package api

import (
	"go-countries-rest-api/api/store"
	server "go-countries-rest-api/api/server"
	"net/http"
)

type App struct {
	Port string
}

func (a *App) Run() {
	mux := http.NewServeMux()
	countriesStorage := store.NewCountriesStorage()
	server := server.Server{
		Mux:     mux,
		Actions: countriesStorage,
	}
	server.Initialize(a.Port)
}
