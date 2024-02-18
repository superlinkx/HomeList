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

func (h Handlers) GetItems(w http.ResponseWriter, r *http.Request, listID int64, params oapiserver.GetItemsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h Handlers) CreateItem(w http.ResponseWriter, r *http.Request, listID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h Handlers) DeleteItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h Handlers) GetItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h Handlers) UpdateItem(w http.ResponseWriter, r *http.Request, listID int64, itemID int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
