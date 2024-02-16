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
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/superlinkx/HomeList/handler"
	"github.com/superlinkx/HomeList/oapiserver"
)

type Config struct {
	HostURL string
}

func NewServer(config Config, hdls handler.Handlers) (*http.Server, error) {
	if config.HostURL == "" {
		config.HostURL = ":80"
	}

	r := chi.NewRouter()

	swagger, err := oapiserver.GetSwagger()
	if err != nil {
		return nil, err
	}

	r.Use(chimiddleware.Logger, oapimiddleware.OapiRequestValidator(swagger))
	r.Mount("/api/v1", v1ApiRouter(hdls))

	return &http.Server{
		Addr:    config.HostURL,
		Handler: r,
	}, nil
}

func v1ApiRouter(hdls handler.Handlers) http.Handler {
	swagger, err := oapiserver.GetSwagger()
	if err != nil {
		log.Fatalf("Failed to get swagger: %v", err)
	}

	r := chi.NewRouter()
	r.Use(oapimiddleware.OapiRequestValidator(swagger))

	return oapiserver.HandlerFromMux(hdls, r)
}
