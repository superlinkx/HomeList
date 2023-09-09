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

func (s API) FetchListItem(w http.ResponseWriter, r *http.Request) {
	var itemID = chi.URLParam(r, "id")

	if listID, err := strconv.Atoi(itemID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listitem, err := s.application.FetchListItem(r.Context(), int64(listID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listitem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
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

func (s API) CreateListItem(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    CreateListItemParams
		listIDParam = chi.URLParam(r, "listID")
	)

	if listID, err := strconv.Atoi(listIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.application.AddItemToList(r.Context(), int64(listID), listItem.Content, listItem.Sort); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s API) UpdateListItemContent(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemContentParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.application.UpdateListItemContent(r.Context(), int64(itemID), listItem.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s API) UpdateListItemSort(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemSortParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.application.UpdateListItemSort(r.Context(), int64(itemID), listItem.Sort); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s API) UpdateListItemChecked(w http.ResponseWriter, r *http.Request) {
	var (
		listItem    UpdateListItemCheckedParams
		itemIDParam = chi.URLParam(r, "id")
	)

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := json.NewDecoder(r.Body).Decode(&listItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if listItem, err := s.application.UpdateListItemChecked(r.Context(), int64(itemID), listItem.Checked); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if result, err := json.Marshal(ListItem(listItem)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func (s API) DeleteListItem(w http.ResponseWriter, r *http.Request) {
	var itemIDParam = chi.URLParam(r, "id")

	if itemID, err := strconv.Atoi(itemIDParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if err := s.application.DeleteListItem(r.Context(), int64(itemID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func listItemsView(listitems []app.ListItem) ([]byte, error) {
	var view = make([]ListItem, 0, len(listitems))
	for _, listitem := range listitems {
		item := ListItem(listitem)
		view = append(view, item)
	}
	return json.Marshal(view)
}
