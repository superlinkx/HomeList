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

type List struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s Handlers) GetLists(w http.ResponseWriter, r *http.Request, params oapiserver.GetListsParams) {
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

	if ls, err := s.app.AllLists(r.Context(), limit, offset); errors.Is(err, app.ErrNotFound) {
		http.Error(w, "No lists found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		var lists = make([]List, 0, len(ls))
		for _, l := range ls {
			lists = append(lists, List(l))
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(lists)
	}
}

func (s Handlers) CreateList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) DeleteList(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) GetList(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) UpdateList(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
