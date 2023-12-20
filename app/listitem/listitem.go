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

	"github.com/superlinkx/HomeList/service/listitem"
)

type ListItemService interface {
	FetchListItem(ctx context.Context, id int64) (listitem.ListItem, error)
	FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]listitem.ListItem, error)
	AddItemToList(ctx context.Context, listID int64, name string, sort int64) (listitem.ListItem, error)
	UpdateListItemContent(ctx context.Context, id int64, content string) (listitem.ListItem, error)
	UpdateListItemSort(ctx context.Context, id int64, sort int64) (listitem.ListItem, error)
	UpdateListItemChecked(ctx context.Context, id int64, checked bool) (listitem.ListItem, error)
	DeleteListItem(ctx context.Context, id int64) error
}

type ListItem struct {
	ID      int64
	ListID  int64
	Content string
	Checked bool
	Sort    int64
}

type ListItemApp struct {
	listItemService ListItemService
}

func NewListItemApp(listItemService ListItemService) ListItemApp {
	return ListItemApp{
		listItemService: listItemService,
	}
}

func (s ListItemApp) FetchListItem(ctx context.Context, itemID int64) (ListItem, error) {
	if result, err := s.listItemService.FetchListItem(ctx, itemID); err != nil {
		return ListItem{}, fmt.Errorf("application failed to fetch list item: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (s ListItemApp) FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]ListItem, error) {
	if results, err := s.listItemService.FetchAllItemsFromList(ctx, listID, limit); err != nil {
		return []ListItem{}, fmt.Errorf("application failed to fetch all items from list with id %d: %w", listID, err)
	} else {
		var listitems = make([]ListItem, 0, len(results))
		for _, result := range results {
			listitem := ListItem{
				ID:      result.ID,
				ListID:  result.ListID,
				Content: result.Content,
				Checked: result.Checked,
				Sort:    result.Sort,
			}
			listitems = append(listitems, listitem)
		}
		return listitems, nil
	}
}

func (s ListItemApp) AddItemToList(ctx context.Context, listID int64, content string, sort int64) (ListItem, error) {
	if result, err := s.listItemService.AddItemToList(ctx, listID, content, sort); err != nil {
		return ListItem{}, fmt.Errorf("application failed to add item to list: %w", err)
	} else {
		return ListItem{
			ID:      result.ID,
			ListID:  result.ListID,
			Content: result.Content,
			Checked: result.Checked,
			Sort:    result.Sort,
		}, nil
	}
}

func (s ListItemApp) UpdateListItemContent(ctx context.Context, itemID int64, content string) (ListItem, error) {
	if result, err := s.listItemService.UpdateListItemContent(ctx, itemID, content); err != nil {
		return ListItem{}, fmt.Errorf("application failed to update list item text: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (s ListItemApp) UpdateListItemChecked(ctx context.Context, itemID int64, checked bool) (ListItem, error) {
	if result, err := s.listItemService.UpdateListItemChecked(ctx, itemID, checked); err != nil {
		return ListItem{}, fmt.Errorf("application failed to update list item checked: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (s ListItemApp) UpdateListItemSort(ctx context.Context, itemID int64, sort int64) (ListItem, error) {
	if result, err := s.listItemService.UpdateListItemSort(ctx, itemID, sort); err != nil {
		return ListItem{}, fmt.Errorf("application failed to update list item sort: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (s ListItemApp) DeleteListItem(ctx context.Context, itemID int64) error {
	if err := s.listItemService.DeleteListItem(ctx, itemID); err != nil {
		return fmt.Errorf("application failed to delete list item: %w", err)
	} else {
		return nil
	}
}
