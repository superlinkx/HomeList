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

package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/superlinkx/HomeList/app/model"
	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/data/sqlite/sqlc"
)

func (s SqliteAdapter) AllItemsFromListPaginated(ctx context.Context, listID int64, limit int32, offset int32) ([]model.Item, error) {
	var params = sqlc.AllItemsFromListPaginatedParams{
		ListID: listID,
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	if items, err := s.queries.AllItemsFromListPaginated(ctx, params); errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("no items found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get items: %w", err)
	} else {
		var adaptedItems = make([]model.Item, 0, len(items))
		for _, item := range items {
			adaptedItems = append(adaptedItems, model.Item(item))
		}
		return adaptedItems, nil
	}
}

func (s SqliteAdapter) CreateItemOnList(ctx context.Context, listID int64, content string, sort int64) (model.Item, error) {
	var params = sqlc.CreateListItemParams{
		ListID:  listID,
		Content: content,
		Sort:    sort,
	}

	if item, err := s.queries.CreateListItem(ctx, params); err != nil {
		return model.Item{}, fmt.Errorf("failed to create item: %w", err)
	} else {
		return model.Item(item), nil
	}
}

func (s SqliteAdapter) GetItemFromList(ctx context.Context, listID int64, id int64) (model.Item, error) {
	if item, err := s.queries.GetListItem(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return model.Item{}, fmt.Errorf("no item found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.Item{}, fmt.Errorf("failed to get item: %w", err)
	} else if item.ListID != listID {
		return model.Item{}, fmt.Errorf("item %d not found in list %d: %w", id, listID, adapter.ErrNotFound)
	} else {
		return model.Item(item), nil
	}
}

func (s SqliteAdapter) UpdateItemFromListContent(ctx context.Context, listID int64, id int64, content string) (model.Item, error) {
	if item, err := s.queries.GetListItem(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return model.Item{}, fmt.Errorf("no item found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.Item{}, fmt.Errorf("failed to get item: %w", err)
	} else if item.ListID != listID {
		return model.Item{}, fmt.Errorf("item %d not found in list %d: %w", id, listID, adapter.ErrNotFound)
	} else if item, err := s.queries.UpdateListItemContent(ctx, sqlc.UpdateListItemContentParams{ID: id, Content: content}); err != nil {
		return model.Item{}, fmt.Errorf("failed to update item: %w", err)
	} else {
		return model.Item(item), nil
	}
}

func (s SqliteAdapter) UpdateItemFromListChecked(ctx context.Context, listID int64, id int64, checked bool) (model.Item, error) {
	if item, err := s.queries.GetListItem(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return model.Item{}, fmt.Errorf("no item found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.Item{}, fmt.Errorf("failed to get item: %w", err)
	} else if item.ListID != listID {
		return model.Item{}, fmt.Errorf("item %d not found in list %d: %w", id, listID, adapter.ErrNotFound)
	} else if item, err := s.queries.UpdateListItemChecked(ctx, sqlc.UpdateListItemCheckedParams{ID: id, Checked: checked}); err != nil {
		return model.Item{}, fmt.Errorf("failed to update item: %w", err)
	} else {
		return model.Item(item), nil
	}
}

func (s SqliteAdapter) UpdateItemFromListSort(ctx context.Context, listID int64, id int64, sort int64) (model.Item, error) {
	if item, err := s.queries.GetListItem(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return model.Item{}, fmt.Errorf("no item found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.Item{}, fmt.Errorf("failed to get item: %w", err)
	} else if item.ListID != listID {
		return model.Item{}, fmt.Errorf("item %d not found in list %d: %w", id, listID, adapter.ErrNotFound)
	} else if item, err := s.queries.UpdateListItemSort(ctx, sqlc.UpdateListItemSortParams{ID: id, Sort: sort}); err != nil {
		return model.Item{}, fmt.Errorf("failed to update item: %w", err)
	} else {
		return model.Item(item), nil
	}
}

func (s SqliteAdapter) DeleteItemFromList(ctx context.Context, listID int64, id int64) error {
	if item, err := s.queries.GetListItem(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("no items found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return fmt.Errorf("failed to get item: %w", err)
	} else if item.ListID != listID {
		return fmt.Errorf("item %d not found in list %d: %w", id, listID, adapter.ErrNotFound)
	} else if err := s.queries.DeleteListItem(ctx, id); err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	} else {
		return nil
	}
}
