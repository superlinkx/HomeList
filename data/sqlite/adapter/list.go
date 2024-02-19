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
	"github.com/superlinkx/HomeList/data/mapper"
	"github.com/superlinkx/HomeList/data/sqlite/sqlc"
)

func (s SqliteAdapter) AllLists(ctx context.Context, limit int32, offset int32) ([]model.List, error) {
	var params = sqlc.AllListsParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	if lists, err := s.queries.AllLists(ctx, params); errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("no lists found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get lists: %w", err)
	} else {
		var adaptedLists = make([]model.List, 0, len(lists))
		for _, list := range lists {
			adaptedLists = append(adaptedLists, model.List(list))
		}
		return adaptedLists, nil
	}
}

func (s SqliteAdapter) GetList(ctx context.Context, id int64) (model.List, error) {
	if list, err := s.queries.GetList(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return model.List{}, fmt.Errorf("no lists found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.List{}, fmt.Errorf("failed to get list: %w", err)
	} else {
		return model.List(list), nil
	}
}

func (s SqliteAdapter) RenameList(ctx context.Context, id int64, name string) (model.List, error) {
	if list, err := s.queries.RenameList(ctx, sqlc.RenameListParams{ID: id, Name: name}); errors.Is(err, sql.ErrNoRows) {
		return model.List{}, fmt.Errorf("no lists found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.List{}, fmt.Errorf("failed to rename list: %w", err)
	} else {
		return model.List(list), nil
	}
}

func (s SqliteAdapter) CreateList(ctx context.Context, name string) (model.List, error) {
	if list, err := s.queries.CreateList(ctx, name); errors.Is(err, sql.ErrNoRows) {
		return model.List{}, fmt.Errorf("no lists found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return model.List{}, fmt.Errorf("failed to create list: %w", err)
	} else {
		return model.List(list), nil
	}
}

func (s SqliteAdapter) DeleteList(ctx context.Context, id int64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	if list, err := qtx.GetList(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("could not find list: %w", adapter.ErrNotFound)
	} else if err != nil {
		return fmt.Errorf("failed to get list: %w", err)
	} else if err := s.deleteAllListItems(ctx, id, qtx); err != nil {
		return fmt.Errorf("failed to delete all items from list: %w", err)
	} else if err := qtx.DeleteList(ctx, list.ID); err != nil {
		return fmt.Errorf("failed to delete list: %w", err)
	} else {
		return tx.Commit()
	}
}

func (s SqliteAdapter) deleteAllListItems(ctx context.Context, listID int64, qtx *sqlc.Queries) error {
	if items, err := qtx.AllItemsFromList(ctx, listID); errors.Is(err, sql.ErrNoRows) {
		// No items to delete, so proceed normally
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get items: %w", err)
	} else {
		for _, item := range items {
			if err := qtx.DeleteListItem(ctx, item.ID); err != nil {
				return fmt.Errorf("failed to delete item %d: %w", item.ID, err)
			}
		}
		return nil
	}
}

func (s SqliteAdapter) ReflowList(ctx context.Context, id int64, reflowMapper mapper.Reflow) error {
	// Start a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	// Make sure we're operating on a real list
	if _, err := qtx.GetList(ctx, id); errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("list not found: %w", adapter.ErrNotFound)
	} else if err != nil {
		return fmt.Errorf("failed to get list: %w", err)
	}

	// Get all items from list
	items, err := qtx.AllItemsFromList(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		// No items found, so reflow is a no-op
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get list items: %w", err)
	}

	// Convert to app model before sending to outside function
	var modelItems = make([]model.Item, 0, len(items))
	for _, item := range items {
		modelItems = append(modelItems, model.Item(item))
	}

	// Reflow the items with the delegate function
	reflowMapper(modelItems)

	// Update each item in the database
	// TODO: Turn this into a batched update
	for _, item := range modelItems {
		if _, err := qtx.UpdateListItemSort(ctx, sqlc.UpdateListItemSortParams{ID: item.ID, Sort: item.Sort}); err != nil {
			return fmt.Errorf("failed to update item %d: %w", item.ID, err)
		}
	}

	return tx.Commit()
}
