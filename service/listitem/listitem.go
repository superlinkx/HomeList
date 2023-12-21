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
	"errors"
	"fmt"

	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/service/common"
)

type queries interface {
	GetListItem(ctx context.Context, id int64) (adapter.ListItem, error)
	AllItemsFromList(ctx context.Context, listID int64, limit int32) ([]adapter.ListItem, error)
	CreateListItem(ctx context.Context, listID int64, content string, sort int64) (adapter.ListItem, error)
	UpdateListItemText(ctx context.Context, id int64, content string) (adapter.ListItem, error)
	UpdateListItemSort(ctx context.Context, id int64, sort int64) (adapter.ListItem, error)
	UpdateListItemChecked(ctx context.Context, id int64, checked bool) (adapter.ListItem, error)
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
	if listitem, err := s.queries.GetListItem(ctx, id); errors.Is(err, adapter.ErrNotFound) {
		return ListItem{}, fmt.Errorf("no list item found: %w", common.ErrNotFound)
	} else if err != nil {
		return ListItem{}, fmt.Errorf("failed to fetch list item: %w", err)
	} else {
		return ListItem(listitem), nil
	}
}

func (s ListItemService) FetchAllItemsFromList(ctx context.Context, listID int64, limit int32) ([]ListItem, error) {
	if results, err := s.queries.AllItemsFromList(ctx, listID, limit); err != nil {
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
	if listItem, err := s.queries.CreateListItem(ctx, listID, content, sort); err != nil {
		return ListItem{}, fmt.Errorf("failed to create list item: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) UpdateListItemContent(ctx context.Context, itemID int64, content string) (ListItem, error) {
	if listItem, err := s.queries.UpdateListItemText(ctx, itemID, content); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item text: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) UpdateListItemSort(ctx context.Context, itemID int64, sort int64) (ListItem, error) {
	if listItem, err := s.queries.UpdateListItemSort(ctx, itemID, sort); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item sort: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s ListItemService) UpdateListItemChecked(ctx context.Context, itemID int64, checked bool) (ListItem, error) {
	if listItem, err := s.queries.UpdateListItemChecked(ctx, itemID, checked); err != nil {
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
