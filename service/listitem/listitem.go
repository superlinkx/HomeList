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
	"fmt"

	"github.com/superlinkx/HomeList/db/sqlite"
)

type queries interface {
	GetListItem(ctx context.Context, id int64) (sqlite.ListItem, error)
	AllItemsFromList(ctx context.Context, params sqlite.AllItemsFromListParams) ([]sqlite.ListItem, error)
	CreateListItem(ctx context.Context, params sqlite.CreateListItemParams) (sqlite.ListItem, error)
	UpdateListItemText(ctx context.Context, params sqlite.UpdateListItemTextParams) (sqlite.ListItem, error)
	UpdateListItemSort(ctx context.Context, params sqlite.UpdateListItemSortParams) (sqlite.ListItem, error)
	UpdateListItemChecked(ctx context.Context, params sqlite.UpdateListItemCheckedParams) (sqlite.ListItem, error)
	DeleteListItem(ctx context.Context, id int64) error
}

type ListItemService struct {
	queries queries
}

type ListItem struct {
	ID      int64
	ListID  int64
	Content string
	Checked bool
	Sort    int64
}

func NewService(q queries) ListItemService {
	return ListItemService{queries: q}
}

func (s ListItemService) FetchListItem(ctx context.Context, id int64) (ListItem, error) {
	if listitem, err := s.queries.GetListItem(ctx, id); err != nil {
		return ListItem{}, fmt.Errorf("failed to fetch list item: %w", err)
	} else {
		return ListItem(listitem), nil
	}
}

func (s ListItemService) FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]ListItem, error) {
	var params = sqlite.AllItemsFromListParams{ListID: listID, Limit: limit}

	if results, err := s.queries.AllItemsFromList(ctx, params); err != nil {
		return []ListItem{}, fmt.Errorf("failed to fetch list items: %w", err)
	} else {
		var listitems = make([]ListItem, 0, len(results))
		for _, result := range results {
			listitem := ListItem(result)
			listitems = append(listitems, listitem)
		}
		return listitems, nil
	}
}

func (s ListItemService) AddItemToList(ctx context.Context, listID int64, content string, sort int64) (ListItem, error) {
	params := sqlite.CreateListItemParams{
		ListID:  listID,
		Content: content,
		Sort:    sort,
	}

	if listItem, err := s.queries.CreateListItem(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to create list item: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) UpdateListItemContent(ctx context.Context, itemID int64, content string) (ListItem, error) {
	params := sqlite.UpdateListItemTextParams{
		ID:      itemID,
		Content: content,
	}

	if listItem, err := s.queries.UpdateListItemText(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item text: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) UpdateListItemSort(ctx context.Context, itemID int64, sort int64) (ListItem, error) {
	params := sqlite.UpdateListItemSortParams{
		ID:   itemID,
		Sort: sort,
	}

	if listItem, err := s.queries.UpdateListItemSort(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item sort: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) UpdateListItemChecked(ctx context.Context, itemID int64, checked bool) (ListItem, error) {
	params := sqlite.UpdateListItemCheckedParams{
		ID:      itemID,
		Checked: checked,
	}

	if listItem, err := s.queries.UpdateListItemChecked(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item checked state: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) DeleteListItem(ctx context.Context, itemID int64) error {
	if err := s.queries.DeleteListItem(ctx, itemID); err != nil {
		return fmt.Errorf("failed to delete list item: %w", err)
	} else {
		return nil
	}
}
