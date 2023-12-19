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

package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ridge/must"
	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/handler"
	"github.com/superlinkx/HomeList/handler/mocks"
)

var (
	appListItem = app.ListItem{
		ID:      1,
		ListID:  1,
		Content: "Test Item",
		Checked: true,
		Sort:    1,
	}
	appListItemSlice = []app.ListItem{
		appListItem,
	}
	handlerListItem      = handler.ListItem(appListItem)
	handlerListItemSlice = []handler.ListItem{
		handlerListItem,
	}
	jsonListItem      = string(must.Bytes(json.Marshal(handlerListItem)))
	jsonListItemSlice = string(must.Bytes(json.Marshal(handlerListItemSlice)))
)

func TestHandlers_FetchListItem(t *testing.T) {
	var (
		mockApplication = mocks.NewMockApplication(t)
		hdls            = handler.NewHandlersWithApplication(mockApplication)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/listitem/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockApplication.EXPECT().FetchListItem(ctx, int64(1)).
			Return(appListItem, nil).Times(1)

		hdls.FetchListItem(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItem, res.Body.String())
	})

	t.Run("bad id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/listitem/invalidkey", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "invalidkey")
		req = req.WithContext(ctx)

		hdls.FetchListItem(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/listitem/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockApplication.EXPECT().FetchListItem(ctx, int64(1)).
			Return(app.ListItem{}, errAppFailure).Times(1)

		hdls.FetchListItem(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestHandlers_FetchAllItemsFromList(t *testing.T) {
	var (
		mockApplication = mocks.NewMockApplication(t)
		hdls            = handler.NewHandlersWithApplication(mockApplication)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/list/1/items", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockApplication.EXPECT().FetchAllItemsFromList(ctx, int64(1), int64(10)).
			Return(appListItemSlice, nil).Times(1)

		hdls.FetchAllItemsFromList(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItemSlice, res.Body.String())
	})
}
