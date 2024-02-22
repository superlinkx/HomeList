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

package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/oapiserver"
)

type Item struct {
	ID      int64  `json:"id"`
	ListID  int64  `json:"list_id"`
	Content string `json:"content"`
	Checked bool   `json:"checked"`
	Sort    int64  `json:"sort"`
}

func (s Handlers) GetItems(w http.ResponseWriter, r *http.Request, listID int64, params oapiserver.GetItemsParams) {
	var (
		limit  int32 = 10
		offset int32 = 0
	)

	if params.Limit != nil {
		limit = *params.Limit
	}

	if params.Offset != nil {
		offset = *params.Offset
	}

	if is, err := s.app.AllItemsFromList(r.Context(), listID, limit, offset); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "list items not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		var items = make([]Item, 0, len(is))
		for _, i := range is {
			items = append(items, Item(i))
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(items)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
		}
	}
}

func (s Handlers) CreateItem(w http.ResponseWriter, r *http.Request, listID int64) {
	var item oapiserver.CreateItemJSONBody
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid body")
	} else if item.Content == "" {
		errorResponse(w, http.StatusBadRequest, "content cannot be empty")
	} else if i, err := s.app.CreateItemOnList(r.Context(), listID, item.Content, item.Sort); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "list not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(i)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
		}
	}
}

func (s Handlers) GetItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	if i, err := s.app.GetItemFromList(r.Context(), listID, itemID); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "item not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(i)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
		}
	}
}

func (s Handlers) UpdateItemFromListContent(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	var item oapiserver.UpdateItemFromListContentJSONBody
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid body")
	} else if item.Content == "" {
		errorResponse(w, http.StatusBadRequest, "content cannot be empty")
	} else if i, err := s.app.UpdateItemFromListContent(r.Context(), listID, itemID, item.Content); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "item not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(i)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
		}
	}
}

func (s Handlers) UpdateItemFromListChecked(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	var item oapiserver.UpdateItemFromListCheckedJSONBody
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid body")
	} else if i, err := s.app.UpdateItemFromListChecked(r.Context(), listID, itemID, item.Checked); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "item not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(i)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
		}
	}
}

func (s Handlers) UpdateItemFromListSort(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	var item oapiserver.UpdateItemFromListSortJSONBody
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		errorResponse(w, http.StatusBadRequest, "invalid body")
	} else if i, err := s.app.UpdateItemFromListSort(r.Context(), listID, itemID, item.Sort); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "item not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(i)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError)
		}
	}
}

func (s Handlers) DeleteItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	if err := s.app.DeleteItemFromList(r.Context(), listID, itemID); errors.Is(err, app.ErrNotFound) {
		errorResponse(w, http.StatusNotFound, "item not found")
	} else if err != nil {
		errorResponse(w, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
