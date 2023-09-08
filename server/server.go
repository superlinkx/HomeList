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
	r.Put("/list/{listID}", api.RenameList)
	r.Delete("/list/{listID}", api.DeleteList)
	return r
}
