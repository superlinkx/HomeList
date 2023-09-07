package api

import (
	"encoding/json"
	"net/http"

	"github.com/superlinkx/HomeList/app"
)

type API struct {
	application app.Application
}

type ViewList struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewAPI(application app.Application) API {
	return API{
		application: application,
	}
}

func (s API) FetchAllLists(w http.ResponseWriter, r *http.Request) {
	if lists, err := s.application.AllLists(r.Context(), 10); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listsView(lists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

// func listView(list app.List) ([]byte, error) {
// 	return json.Marshal(ViewList{ID: list.ID, Name: list.Name})
// }

func listsView(lists []app.List) ([]byte, error) {
	var viewLists = make([]ViewList, 0, len(lists))
	for _, list := range lists {
		viewLists = append(viewLists, ViewList{ID: list.ID, Name: list.Name})
	}
	return json.Marshal(viewLists)
}
