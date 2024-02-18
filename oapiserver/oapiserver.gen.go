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

// Package oapiserver provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package oapiserver

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// GetListsParams defines parameters for GetLists.
type GetListsParams struct {
	Limit  *int32 `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`
}

// CreateListJSONBody defines parameters for CreateList.
type CreateListJSONBody struct {
	Name string `json:"name"`
}

// UpdateListJSONBody defines parameters for UpdateList.
type UpdateListJSONBody struct {
	Name string `json:"name"`
}

// GetItemsParams defines parameters for GetItems.
type GetItemsParams struct {
	Limit  *int32 `form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`
}

// CreateItemJSONBody defines parameters for CreateItem.
type CreateItemJSONBody struct {
	Content string `json:"content"`
	Sort    int64  `json:"sort"`
}

// UpdateItemJSONBody defines parameters for UpdateItem.
type UpdateItemJSONBody struct {
	Name string `json:"name"`
}

// CreateListJSONRequestBody defines body for CreateList for application/json ContentType.
type CreateListJSONRequestBody CreateListJSONBody

// UpdateListJSONRequestBody defines body for UpdateList for application/json ContentType.
type UpdateListJSONRequestBody UpdateListJSONBody

// CreateItemJSONRequestBody defines body for CreateItem for application/json ContentType.
type CreateItemJSONRequestBody CreateItemJSONBody

// UpdateItemJSONRequestBody defines body for UpdateItem for application/json ContentType.
type UpdateItemJSONRequestBody UpdateItemJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of lists
	// (GET /lists)
	GetLists(w http.ResponseWriter, r *http.Request, params GetListsParams)
	// Create a new list
	// (POST /lists)
	CreateList(w http.ResponseWriter, r *http.Request)
	// Delete list by ID
	// (DELETE /lists/{listID})
	DeleteList(w http.ResponseWriter, r *http.Request, listID int64)
	// Get list by ID
	// (GET /lists/{listID})
	GetList(w http.ResponseWriter, r *http.Request, listID int64)
	// Update list by ID
	// (PUT /lists/{listID})
	UpdateList(w http.ResponseWriter, r *http.Request, listID int64)
	// Get a list of items from listID
	// (GET /lists/{listID}/items)
	GetItems(w http.ResponseWriter, r *http.Request, listID int64, params GetItemsParams)
	// Create a new item on the listID
	// (POST /lists/{listID}/items)
	CreateItem(w http.ResponseWriter, r *http.Request, listID int64)
	// Delete item by listID and itemID
	// (DELETE /lists/{listID}/items/{itemID})
	DeleteItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64)
	// Get item by listID and itemID
	// (GET /lists/{listID}/items/{itemID})
	GetItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64)
	// Update item by listID and itemID
	// (PUT /lists/{listID}/items/{itemID})
	UpdateItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get a list of lists
// (GET /lists)
func (_ Unimplemented) GetLists(w http.ResponseWriter, r *http.Request, params GetListsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new list
// (POST /lists)
func (_ Unimplemented) CreateList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete list by ID
// (DELETE /lists/{listID})
func (_ Unimplemented) DeleteList(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get list by ID
// (GET /lists/{listID})
func (_ Unimplemented) GetList(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update list by ID
// (PUT /lists/{listID})
func (_ Unimplemented) UpdateList(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a list of items from listID
// (GET /lists/{listID}/items)
func (_ Unimplemented) GetItems(w http.ResponseWriter, r *http.Request, listID int64, params GetItemsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new item on the listID
// (POST /lists/{listID}/items)
func (_ Unimplemented) CreateItem(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete item by listID and itemID
// (DELETE /lists/{listID}/items/{itemID})
func (_ Unimplemented) DeleteItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get item by listID and itemID
// (GET /lists/{listID}/items/{itemID})
func (_ Unimplemented) GetItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update item by listID and itemID
// (PUT /lists/{listID}/items/{itemID})
func (_ Unimplemented) UpdateItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetLists operation middleware
func (siw *ServerInterfaceWrapper) GetLists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetListsParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLists(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateList operation middleware
func (siw *ServerInterfaceWrapper) CreateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateList(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteList operation middleware
func (siw *ServerInterfaceWrapper) DeleteList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteList(w, r, listID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetList operation middleware
func (siw *ServerInterfaceWrapper) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetList(w, r, listID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateList operation middleware
func (siw *ServerInterfaceWrapper) UpdateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateList(w, r, listID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetItems operation middleware
func (siw *ServerInterfaceWrapper) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetItemsParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItems(w, r, listID, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateItem operation middleware
func (siw *ServerInterfaceWrapper) CreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateItem(w, r, listID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteItem operation middleware
func (siw *ServerInterfaceWrapper) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	// ------------- Path parameter "itemID" -------------
	var itemID int64

	err = runtime.BindStyledParameterWithOptions("simple", "itemID", chi.URLParam(r, "itemID"), &itemID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteItem(w, r, listID, itemID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetItem operation middleware
func (siw *ServerInterfaceWrapper) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	// ------------- Path parameter "itemID" -------------
	var itemID int64

	err = runtime.BindStyledParameterWithOptions("simple", "itemID", chi.URLParam(r, "itemID"), &itemID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItem(w, r, listID, itemID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateItem operation middleware
func (siw *ServerInterfaceWrapper) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "listID" -------------
	var listID int64

	err = runtime.BindStyledParameterWithOptions("simple", "listID", chi.URLParam(r, "listID"), &listID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "listID", Err: err})
		return
	}

	// ------------- Path parameter "itemID" -------------
	var itemID int64

	err = runtime.BindStyledParameterWithOptions("simple", "itemID", chi.URLParam(r, "itemID"), &itemID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateItem(w, r, listID, itemID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/lists", wrapper.GetLists)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/lists", wrapper.CreateList)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/lists/{listID}", wrapper.DeleteList)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/lists/{listID}", wrapper.GetList)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/lists/{listID}", wrapper.UpdateList)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/lists/{listID}/items", wrapper.GetItems)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/lists/{listID}/items", wrapper.CreateItem)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/lists/{listID}/items/{itemID}", wrapper.DeleteItem)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/lists/{listID}/items/{itemID}", wrapper.GetItem)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/lists/{listID}/items/{itemID}", wrapper.UpdateItem)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RY34/iNhD+V6xpH1opWsMe24e8tV21RdpWfWifTqvKJAP4mtg+26FFKP975R8kQEiA",
	"W8pxTwtkPJ75vvlmZrOBTJZKChTWQLoBxTQr0aL237jFkuf+k4AUFLNLSECwEiH1D6fPkIDGjxXXmENq",
	"dYUJmGyJJXOn5lKXzDpbYb+bQAJ2rTB8xQVqqOsECl5y21zxsUK9bu8ID/tcvnvsc2lsb9ju4RXClvO5",
	"wd6449PLAq9DTGjsDzLn2DDg/mZSWBT+PqZUwTNmuRT0g5HC/dZeo7RUqG08vnMsXmes5mIBdQJGantu",
	"ti1W7xuf0cNrc0DOPmBmQyI5mkxz5YKEFF64sWQm8zWxksyQZBqZxbxDQuTuDQkH/DvZHqTgra4auL/A",
	"KClMiGMyGh3ij/9aqgrGD+I/jLQTw1SsWMFz8vtWmt+Ybx1/k9HkWlf8Ji35SVYid36frhm6RS1YQQzq",
	"FWqCWstQUMHJXokfVO4Ss78x37lmJmWBTLgQd2IruXhBsbBLSMdJt8hDGzhZ4qHu/jrb+u3a2ebXq6JW",
	"CvvAnB3jVgqDCJ0vDC7m0rkreIbCYCs2+HX6ByRQ6QJSWFqrTEqpVCiMrHSGD1IvaDxkqLOtE7DcFu7o",
	"L7JEpzFIYIXahKoZOwvngCkOKbx7GD+MIPFt3CNAHTD+0yJ0YAePbw/THFL4Ge2LN0j2ptn7DXytcQ4p",
	"fEXbmUdbExqGTZ2cNIzdvX49EP1jRznDrctVvj947L6oEJ+shyxQwrRm62Na+544SyLnJMDTdohj3pu4",
	"qTO6meqrsmR6HUgirBuxkuYIoz/6rhvrpB2R6/7kdqZoA+ABVeOLqDrNUM/w2E4MT8foHDpGO3QM2zqj",
	"fVgDUoQRgf+QGFcSBUM3YfWpneMcC7TYhfrZ/x6hvlQ+fus6oopJuLGDTQjiYmw+oawvxTHAEOpztibT",
	"Z+drqNtcEa3R/16YTnz7qd0Z/N0AVXUE+z9Vzq5RrNdpKKPbNJTKJ32Pogl07BHXbT60GXp9cpp6g08l",
	"NPlCx7zfgs8Y835NbwZncP0ljXofMZlrWZL4n/iJse/K4aby3jJx833BZfo59wWXOJGC2CU23PQJmG7C",
	"658zlok3EXhapfEd1Xlrh4f47tcOT8RsHUkgTOQkvmwb2EI+G8zXm3tBeT3CmG/fT9zhrjJI2MDqckvO",
	"7m3JGST77pecAca9vR+Hgc72xUhKaSEzViylsenj09MjZYrT1Rjq1/q/AAAA//8S0qxEiRcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
