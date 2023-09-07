package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/superlinkx/HomeList/ui"
)

type Config struct {
	HostURL string
}

func StartServer(config Config, ui ui.UI) {
	if config.HostURL == "" {
		config.HostURL = ":80"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ui/lists", ui.FetchAllLists)
	log.Printf("Server is running on %s", config.HostURL)
	http.ListenAndServe(config.HostURL, r)
}
