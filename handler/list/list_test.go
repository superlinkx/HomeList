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

package list_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ridge/must"
	"github.com/stretchr/testify/assert"
	listapp "github.com/superlinkx/HomeList/app/list"
	"github.com/superlinkx/HomeList/handler/list"
	"github.com/superlinkx/HomeList/handler/list/mocks"
)

var (
	errAppFailure = errors.New("application failed")
	appList       = listapp.List{
		ID:   1,
		Name: "Test List",
	}
	appListSlice = []listapp.List{
		appList,
	}
	handlerList      = list.ListView(appList)
	handlerListSlice = []list.ListView{
		handlerList,
	}
	jsonList      = string(must.Bytes(json.Marshal(handlerList)))
	jsonListSlice = string(must.Bytes(json.Marshal(handlerListSlice)))
	listRequest   = list.ListRequest{
		Name: "Test List",
	}
	jsonListRequest        = must.Bytes(json.Marshal(listRequest))
	jsonInvalidListRequest = []byte(`{"name": []}`)
)

func TestListHandlers_FetchAllLists(t *testing.T) {
	var (
		mockListApp = mocks.NewMockListApp(t)
		hdls        = list.NewHandlers(mockListApp)
	)

	t.Run("empty list", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/lists", nil)
			res = httptest.NewRecorder()
		)

		mockListApp.EXPECT().AllLists(context.Background(), int64(10)).
			Return([]listapp.List{}, nil).Times(1)

		hdls.FetchAllLists(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, "[]", res.Body.String())
	})

	t.Run("happy path list", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/lists", nil)
			res = httptest.NewRecorder()
		)

		mockListApp.EXPECT().AllLists(context.Background(), int64(10)).
			Return(appListSlice, nil).Times(1)

		hdls.FetchAllLists(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListSlice, res.Body.String())
	})

	t.Run("application failed", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodGet, "/lists", nil)
			res = httptest.NewRecorder()
		)

		mockListApp.EXPECT().AllLists(context.Background(), int64(10)).
			Return(nil, errAppFailure).Times(1)

		hdls.FetchAllLists(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListHandlers_FetchList(t *testing.T) {
	var (
		mockListApp = mocks.NewMockListApp(t)
		hdls        = list.NewHandlers(mockListApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/list/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListApp.EXPECT().GetList(ctx, int64(1)).
			Return(appList, nil).Times(1)

		hdls.FetchList(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonList, res.Body.String())
	})

	t.Run("bad request", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/list/invalidkey", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "invalidkey")
		req = req.WithContext(ctx)

		hdls.FetchList(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/list/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListApp.EXPECT().GetList(ctx, int64(1)).
			Return(listapp.List{}, errAppFailure).Times(1)

		hdls.FetchList(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListHandlers_CreateList(t *testing.T) {
	var (
		mockListApp = mocks.NewMockListApp(t)
		hdls        = list.NewHandlers(mockListApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/list", bytes.NewBuffer(jsonListRequest))
			res = httptest.NewRecorder()
		)

		mockListApp.EXPECT().CreateList(context.Background(), listRequest.Name).
			Return(appList, nil).Times(1)

		hdls.CreateList(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonList, res.Body.String())
	})

	t.Run("invalid body", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/list", bytes.NewBuffer(jsonInvalidListRequest))
			res = httptest.NewRecorder()
		)

		hdls.CreateList(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/list", bytes.NewBuffer(jsonListRequest))
			res = httptest.NewRecorder()
		)

		mockListApp.EXPECT().CreateList(context.Background(), listRequest.Name).
			Return(listapp.List{}, errAppFailure).Times(1)

		hdls.CreateList(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListHandlers_RenameList(t *testing.T) {
	var (
		mockListApp = mocks.NewMockListApp(t)
		hdls        = list.NewHandlers(mockListApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/list/1", bytes.NewBuffer(jsonListRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListApp.EXPECT().UpdateList(ctx, int64(1), listRequest.Name).
			Return(appList, nil).Times(1)

		hdls.RenameList(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonList, res.Body.String())
	})

	t.Run("bad list id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/list/invalidkey", bytes.NewBuffer(jsonListRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "invalidkey")
		req = req.WithContext(ctx)

		hdls.RenameList(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("invalid body", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/list/1", bytes.NewBuffer(jsonInvalidListRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		hdls.RenameList(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/list/1", bytes.NewBuffer(jsonListRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListApp.EXPECT().UpdateList(ctx, int64(1), listRequest.Name).
			Return(listapp.List{}, errAppFailure).Times(1)

		hdls.RenameList(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListHandlers_DeleteList(t *testing.T) {
	var (
		mockListApp = mocks.NewMockListApp(t)
		hdls        = list.NewHandlers(mockListApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodDelete, "/list/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListApp.EXPECT().DeleteList(ctx, int64(1)).
			Return(nil).Times(1)

		hdls.DeleteList(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
		assert.Empty(t, res.Body.Bytes())
	})

	t.Run("bad list id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodDelete, "/list/invalidkey", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "invalidkey")
		req = req.WithContext(ctx)

		hdls.DeleteList(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodDelete, "/list/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListApp.EXPECT().DeleteList(ctx, int64(1)).
			Return(errAppFailure).Times(1)

		hdls.DeleteList(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
