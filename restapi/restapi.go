// MIT License
//
// Copyright (c) 2023 Alyx Holms
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package restapi

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/superlinkx/HomeList/handlers"
)

type Config struct {
	HostURL string
}

func NewServer(config Config, api handlers.Handlers) *http.Server {
	if config.HostURL == "" {
		config.HostURL = ":80"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1", v1ApiRouter(api))
	log.Printf("Server is running on %s", config.HostURL)

	return &http.Server{
		Addr:    config.HostURL,
		Handler: r,
	}
}

func v1ApiRouter(api handlers.Handlers) http.Handler {
	r := chi.NewRouter()

	r.Get("/lists", api.FetchAllLists)
	r.Get("/list/{listID}", api.FetchList)
	r.Get("/list/{listID}/items", api.FetchAllItemsFromList)
	r.Get("/listitem/{id}", api.FetchListItem)

	r.Post("/list", api.CreateList)
	r.Post("/list/{listID}/item", api.CreateListItem)

	r.Put("/list/{listID}", api.RenameList)
	r.Put("/listitem/{id}/content", api.UpdateListItemContent)
	r.Put("/listitem/{id}/sort", api.UpdateListItemSort)
	r.Put("/listitem/{id}/checked", api.UpdateListItemChecked)

	r.Delete("/list/{listID}", api.DeleteList)
	r.Delete("/listitem/{id}", api.DeleteListItem)

	return r
}
