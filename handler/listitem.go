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
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superlinkx/HomeList/app"
)

type ListItem struct {
	ID      int64  `json:"id"`
	ListID  int64  `json:"list_id"`
	Content string `json:"content"`
	Checked bool   `json:"checked"`
	Sort    int64  `json:"sort"`
}

type CreateListItemParams struct {
	Content string `json:"content"`
	Sort    int64  `json:"sort"`
}

type UpdateListItemContentParams struct {
	Content string `json:"content"`
}

type UpdateListItemSortParams struct {
	Sort int64 `json:"sort"`
}

type UpdateListItemCheckedParams struct {
	Checked bool `json:"checked"`
}

func (s Handlers) FetchListItem(w http.ResponseWriter, r *http.Request) {
	var itemID = chi.URLParam(r, "id")

	if listID, err := strconv.Atoi(itemID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listitem, err := s.listItemer.FetchListItem(r.Context(), int64(listID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listitem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s Handlers) FetchAllItemsFromList(w http.ResponseWriter, r *http.Request) {
	var listIDParam = chi.URLParam(r, "listID")

	if listID, err := strconv.Atoi(listIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listitems, err := s.listItemer.FetchAllItemsFromList(r.Context(), int64(listID), 10); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listItemsView(listitems); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s Handlers) CreateListItem(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    CreateListItemParams
		listIDParam = chi.URLParam(r, "listID")
	)

	if listID, err := strconv.Atoi(listIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemer.AddItemToList(r.Context(), int64(listID), listItem.Content, listItem.Sort); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s Handlers) UpdateListItemContent(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemContentParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemer.UpdateListItemContent(r.Context(), int64(itemID), listItem.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s Handlers) UpdateListItemSort(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemSortParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemer.UpdateListItemSort(r.Context(), int64(itemID), listItem.Sort); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s Handlers) UpdateListItemChecked(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemCheckedParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemer.UpdateListItemChecked(r.Context(), int64(itemID), listItem.Checked); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s Handlers) DeleteListItem(w http.ResponseWriter, r *http.Request) {
	var itemIDParam = chi.URLParam(r, "id")

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := s.listItemer.DeleteListItem(r.Context(), int64(itemID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func listItemsView(listitems []app.ListItem) ([]byte, error) {
	var view = make([]ListItem, 0, len(listitems))
	for _, listitem := range listitems {
		view = append(view, ListItem(listitem))
	}
	return json.Marshal(view)
}
