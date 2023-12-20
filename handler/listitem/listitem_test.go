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

package listitem_test

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
	listitemapp "github.com/superlinkx/HomeList/app/listitem"
	"github.com/superlinkx/HomeList/handler/listitem"
	"github.com/superlinkx/HomeList/handler/listitem/mocks"
)

var (
	errAppFailure = errors.New("application failed")
	appListItem   = listitemapp.ListItem{
		ID:      1,
		ListID:  1,
		Content: "Test Item",
		Checked: true,
		Sort:    1,
	}
	appListItemSlice = []listitemapp.ListItem{
		appListItem,
	}
	handlerListItem      = listitem.ListItemView(appListItem)
	handlerListItemSlice = []listitem.ListItemView{
		handlerListItem,
	}
	listItemRequest = listitem.CreateListItemParams{
		Content: "Test Item",
		Sort:    1,
	}
	listItemContentRequest = listitem.UpdateListItemContentParams{
		Content: "Test Item",
	}
	listItemSortRequest = listitem.UpdateListItemSortParams{
		Sort: 1,
	}
	listItemCheckedRequest = listitem.UpdateListItemCheckedParams{
		Checked: true,
	}
	jsonListItem                      = string(must.Bytes(json.Marshal(handlerListItem)))
	jsonListItemSlice                 = string(must.Bytes(json.Marshal(handlerListItemSlice)))
	jsonListItemRequest               = must.Bytes(json.Marshal(listItemRequest))
	jsonListItemContentRequest        = must.Bytes(json.Marshal(listItemContentRequest))
	jsonListItemSortRequest           = must.Bytes(json.Marshal(listItemSortRequest))
	jsonListItemCheckedRequest        = must.Bytes(json.Marshal(listItemCheckedRequest))
	jsonInvalidListItemRequest        = []byte(`{"content": []}`)
	jsonInvalidListItemSortRequest    = []byte(`{"sort": []}`)
	jsonInvalidListItemCheckedRequest = []byte(`{"checked": []}`)
)

func TestListItemHandlers_FetchListItem(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
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

		mockListItemApp.EXPECT().FetchListItem(ctx, int64(1)).
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

		mockListItemApp.EXPECT().FetchListItem(ctx, int64(1)).
			Return(listitemapp.ListItem{}, errAppFailure).Times(1)

		hdls.FetchListItem(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListItemHandlers_FetchAllItemsFromList(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
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

		mockListItemApp.EXPECT().FetchAllItemsFromList(ctx, int64(1), int64(10)).
			Return(appListItemSlice, nil).Times(1)

		hdls.FetchAllItemsFromList(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItemSlice, res.Body.String())
	})

	t.Run("invalid list id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/list/invalidkey/items", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "invalidkey")
		req = req.WithContext(ctx)

		hdls.FetchAllItemsFromList(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodGet, "/list/1/items", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().FetchAllItemsFromList(ctx, int64(1), int64(10)).
			Return(nil, errAppFailure).Times(1)

		hdls.FetchAllItemsFromList(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListItemHandlers_CreateListItem(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPost, "/list/1/item", bytes.NewBuffer(jsonListItemRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().AddItemToList(ctx, appListItem.ListID, appListItem.Content, appListItem.Sort).
			Return(appListItem, nil).Times(1)

		hdls.CreateListItem(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItem, res.Body.String())
	})

	t.Run("invalid list id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPost, "/list/invalidkey/item", bytes.NewBuffer(jsonListItemRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "invalidkey")
		req = req.WithContext(ctx)

		hdls.CreateListItem(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("invalid request body", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPost, "/list/1/item", bytes.NewBuffer(jsonInvalidListItemRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		hdls.CreateListItem(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPost, "/list/1/item", bytes.NewBuffer(jsonListItemRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("listID", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().AddItemToList(ctx, appListItem.ListID, appListItem.Content, appListItem.Sort).
			Return(listitemapp.ListItem{}, errAppFailure).Times(1)

		hdls.CreateListItem(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListItemHandlers_UpdateListItemContent(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/content", bytes.NewBuffer(jsonListItemContentRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().UpdateListItemContent(ctx, int64(1), appListItem.Content).
			Return(appListItem, nil).Times(1)

		hdls.UpdateListItemContent(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItem, res.Body.String())
	})

	t.Run("invalid list item id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/invalidkey/content", bytes.NewBuffer(jsonListItemContentRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "invalidkey")
		req = req.WithContext(ctx)

		hdls.UpdateListItemContent(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("invalid request body", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/content", bytes.NewBuffer(jsonInvalidListItemRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		hdls.UpdateListItemContent(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/content", bytes.NewBuffer(jsonListItemContentRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().UpdateListItemContent(ctx, int64(1), appListItem.Content).
			Return(listitemapp.ListItem{}, errAppFailure).Times(1)

		hdls.UpdateListItemContent(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListItemHandlers_UpdateListItemSort(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/sort", bytes.NewBuffer(jsonListItemSortRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().UpdateListItemSort(ctx, int64(1), appListItem.Sort).
			Return(appListItem, nil).Times(1)

		hdls.UpdateListItemSort(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItem, res.Body.String())
	})

	t.Run("invalid list item id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/invalidkey/sort", bytes.NewBuffer(jsonListItemSortRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "invalidkey")
		req = req.WithContext(ctx)

		hdls.UpdateListItemSort(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("invalid request body", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/sort", bytes.NewBuffer(jsonInvalidListItemSortRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		hdls.UpdateListItemSort(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/sort", bytes.NewBuffer(jsonListItemSortRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().UpdateListItemSort(ctx, int64(1), appListItem.Sort).
			Return(listitemapp.ListItem{}, errAppFailure).Times(1)

		hdls.UpdateListItemSort(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListItemHandlers_UpdateListItemChecked(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/checked", bytes.NewBuffer(jsonListItemCheckedRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().UpdateListItemChecked(ctx, int64(1), appListItem.Checked).
			Return(appListItem, nil).Times(1)

		hdls.UpdateListItemChecked(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
		assert.Equal(t, jsonListItem, res.Body.String())
	})

	t.Run("invalid list item id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/invalidkey/checked", bytes.NewBuffer(jsonListItemCheckedRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "invalidkey")
		req = req.WithContext(ctx)

		hdls.UpdateListItemChecked(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("invalid request body", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/checked", bytes.NewBuffer(jsonInvalidListItemCheckedRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		hdls.UpdateListItemChecked(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodPut, "/listitem/1/checked", bytes.NewBuffer(jsonListItemCheckedRequest))
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().UpdateListItemChecked(ctx, int64(1), appListItem.Checked).
			Return(listitemapp.ListItem{}, errAppFailure).Times(1)

		hdls.UpdateListItemChecked(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestListItemHandlers_DeleteListItem(t *testing.T) {
	var (
		mockListItemApp = mocks.NewMockListItemApp(t)
		hdls            = listitem.NewHandlers(mockListItemApp)
	)

	t.Run("happy path", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodDelete, "/listitem/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().DeleteListItem(ctx, int64(1)).
			Return(nil).Times(1)

		hdls.DeleteListItem(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("invalid list item id", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodDelete, "/listitem/invalidkey", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "invalidkey")
		req = req.WithContext(ctx)

		hdls.DeleteListItem(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("application failure", func(t *testing.T) {
		var (
			req  = httptest.NewRequest(http.MethodDelete, "/listitem/1", nil)
			res  = httptest.NewRecorder()
			rctx = chi.NewRouteContext()
			ctx  = context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
		)

		rctx.URLParams.Add("id", "1")
		req = req.WithContext(ctx)

		mockListItemApp.EXPECT().DeleteListItem(ctx, int64(1)).
			Return(errAppFailure).Times(1)

		hdls.DeleteListItem(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
