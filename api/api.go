package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superlinkx/HomeList/app"
)

type API struct {
	application app.Application
}

type ListRequest struct {
	Name string `json:"name"`
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

func (s API) FetchList(w http.ResponseWriter, r *http.Request) {
	var listID = chi.URLParam(r, "listID")

	if id, err := strconv.Atoi(listID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if list, err := s.application.GetList(r.Context(), int64(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listView(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s API) CreateList(w http.ResponseWriter, r *http.Request) {
	var request ListRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if list, err := s.application.CreateList(r.Context(), request.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := listView(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func listView(list app.List) ([]byte, error) {
	return json.Marshal(ViewList{ID: list.ID, Name: list.Name})
}

func listsView(lists []app.List) ([]byte, error) {
	var viewLists = make([]ViewList, 0, len(lists))
	for _, list := range lists {
		viewLists = append(viewLists, ViewList{ID: list.ID, Name: list.Name})
	}
	return json.Marshal(viewLists)
}
