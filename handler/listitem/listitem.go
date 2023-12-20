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

package listitem

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superlinkx/HomeList/app/listitem"
)

type ListItemApp interface {
	FetchAllItemsFromList(ctx context.Context, listID int64, limit int32) ([]listitem.ListItem, error)
	FetchListItem(ctx context.Context, id int64) (listitem.ListItem, error)
	AddItemToList(ctx context.Context, listID int64, content string, sort int64) (listitem.ListItem, error)
	UpdateListItemContent(ctx context.Context, id int64, content string) (listitem.ListItem, error)
	UpdateListItemChecked(ctx context.Context, id int64, checked bool) (listitem.ListItem, error)
	UpdateListItemSort(ctx context.Context, id int64, sort int64) (listitem.ListItem, error)
	DeleteListItem(ctx context.Context, id int64) error
}

type ListItemHandlers struct {
	listItemApp ListItemApp
}

type ListItemView struct {
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

func NewHandlers(li ListItemApp) ListItemHandlers {
	return ListItemHandlers{
		listItemApp: li,
	}
}

func (s ListItemHandlers) FetchListItem(w http.ResponseWriter, r *http.Request) {
	var itemID = chi.URLParam(r, "id")

	if listID, err := strconv.Atoi(itemID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listitem, err := s.listItemApp.FetchListItem(r.Context(), int64(listID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItemView(listitem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListItemHandlers) FetchAllItemsFromList(w http.ResponseWriter, r *http.Request) {
	var listIDParam = chi.URLParam(r, "listID")

	if listID, err := strconv.Atoi(listIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listitems, err := s.listItemApp.FetchAllItemsFromList(r.Context(), int64(listID), 10); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listItemsView(listitems); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListItemHandlers) CreateListItem(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    CreateListItemParams
		listIDParam = chi.URLParam(r, "listID")
	)

	if listID, err := strconv.Atoi(listIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemApp.AddItemToList(r.Context(), int64(listID), listItem.Content, listItem.Sort); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItemView(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListItemHandlers) UpdateListItemContent(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemContentParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemApp.UpdateListItemContent(r.Context(), int64(itemID), listItem.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItemView(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListItemHandlers) UpdateListItemSort(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemSortParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemApp.UpdateListItemSort(r.Context(), int64(itemID), listItem.Sort); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItemView(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListItemHandlers) UpdateListItemChecked(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemCheckedParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.listItemApp.UpdateListItemChecked(r.Context(), int64(itemID), listItem.Checked); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItemView(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListItemHandlers) DeleteListItem(w http.ResponseWriter, r *http.Request) {
	var itemIDParam = chi.URLParam(r, "id")

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := s.listItemApp.DeleteListItem(r.Context(), int64(itemID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func listItemsView(listitems []listitem.ListItem) ([]byte, error) {
	var view = make([]ListItemView, 0, len(listitems))
	for _, listitem := range listitems {
		view = append(view, ListItemView(listitem))
	}
	return json.Marshal(view)
}
