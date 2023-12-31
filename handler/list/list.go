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

package list

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superlinkx/HomeList/app/list"
)

type ListApp interface {
	AllLists(ctx context.Context, limit int32) ([]list.List, error)
	GetList(ctx context.Context, id int64) (list.List, error)
	CreateList(ctx context.Context, name string) (list.List, error)
	UpdateList(ctx context.Context, id int64, name string) (list.List, error)
	DeleteList(ctx context.Context, id int64) error
}

type ListHandlers struct {
	listApp ListApp
}

type ListRequest struct {
	Name string `json:"name"`
}

type ListView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewHandlers(l ListApp) ListHandlers {
	return ListHandlers{
		listApp: l,
	}
}

func (s ListHandlers) FetchAllLists(w http.ResponseWriter, r *http.Request) {
	if lists, err := s.listApp.AllLists(r.Context(), 10); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listsView(lists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListHandlers) FetchList(w http.ResponseWriter, r *http.Request) {
	var listID = chi.URLParam(r, "listID")

	if id, err := strconv.Atoi(listID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if list, err := s.listApp.GetList(r.Context(), int64(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListView(list)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListHandlers) CreateList(w http.ResponseWriter, r *http.Request) {
	var request ListRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if list, err := s.listApp.CreateList(r.Context(), request.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListView(list)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListHandlers) RenameList(w http.ResponseWriter, r *http.Request) {
	var (
		request ListRequest
		listID  = chi.URLParam(r, "listID")
	)

	if id, err := strconv.Atoi(listID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if list, err := s.listApp.UpdateList(r.Context(), int64(id), request.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListView(list)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s ListHandlers) DeleteList(w http.ResponseWriter, r *http.Request) {
	var listID = chi.URLParam(r, "listID")

	if id, err := strconv.Atoi(listID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := s.listApp.DeleteList(r.Context(), int64(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func listsView(lists []list.List) ([]byte, error) {
	var viewLists = make([]ListView, 0, len(lists))
	for _, list := range lists {
		viewLists = append(viewLists, ListView(list))
	}
	return json.Marshal(viewLists)
}
