package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/superlinkx/HomeList/api"
)

type Config struct {
	HostURL string
}

func StartServer(config Config, api api.API) {
	if config.HostURL == "" {
		config.HostURL = ":80"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1", v1ApiRouter(api))
	log.Printf("Server is running on %s", config.HostURL)
	http.ListenAndServe(config.HostURL, r)
}

func v1ApiRouter(api api.API) http.Handler {
	r := chi.NewRouter()
	r.Get("/lists", api.FetchAllLists)
	r.Get("/list/{listID}", api.FetchList)
	r.Post("/list", api.CreateList)
	return r
}
