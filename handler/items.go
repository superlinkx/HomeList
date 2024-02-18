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
	"net/http"

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
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) CreateItem(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) DeleteItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) GetItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Handlers) UpdateItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
