package api

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

func (s API) FetchAllItemsFromList(w http.ResponseWriter, r *http.Request) {
	var listIDParam = chi.URLParam(r, "listID")

	if listID, err := strconv.Atoi(listIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listitems, err := s.application.FetchAllItemsFromList(r.Context(), int64(listID), 10); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listItemsView(listitems); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func listItemsView(listitems []app.ListItem) ([]byte, error) {
	var view = make([]ListItem, 0, len(listitems))
	for _, listitem := range listitems {
		item := ListItem{
			ID:      listitem.ID,
			ListID:  listitem.ListID,
			Content: listitem.Content,
			Checked: listitem.Checked,
			Sort:    listitem.Sort,
		}
		view = append(view, item)
	}
	return json.Marshal(view)
}
