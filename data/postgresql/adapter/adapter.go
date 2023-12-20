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
	"fmt"

	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/data/postgresql/sqlc"
)

type PostgresAdapter struct {
	queries *sqlc.Queries
}

func NewPostgresAdapter(db *sql.DB) adapter.Adapter {
	return PostgresAdapter{queries: sqlc.New(db)}
}

func (s PostgresAdapter) AllLists(ctx context.Context, limit int32) ([]adapter.List, error) {
	if lists, err := s.queries.AllLists(ctx, limit); err != nil {
		return nil, fmt.Errorf("failed to get lists: %w", err)
	} else {
		var adaptedLists = make([]adapter.List, 0, len(lists))
		for _, list := range lists {
			adaptedLists = append(adaptedLists, adapter.List(list))
		}
		return adaptedLists, nil
	}
}

func (s PostgresAdapter) GetList(ctx context.Context, id int64) (adapter.List, error) {
	if list, err := s.queries.GetList(ctx, id); err != nil {
		return adapter.List{}, fmt.Errorf("failed to get list: %w", err)
	} else {
		return adapter.List(list), nil
	}
}

func (s PostgresAdapter) RenameList(ctx context.Context, id int64, name string) (adapter.List, error) {
	if list, err := s.queries.RenameList(ctx, sqlc.RenameListParams{ID: id, Name: name}); err != nil {
		return adapter.List{}, fmt.Errorf("failed to rename list: %w", err)
	} else {
		return adapter.List(list), nil
	}
}

func (s PostgresAdapter) CreateList(ctx context.Context, name string) (adapter.List, error) {
	if list, err := s.queries.CreateList(ctx, name); err != nil {
		return adapter.List{}, fmt.Errorf("failed to create list: %w", err)
	} else {
		return adapter.List(list), nil
	}
}

func (s PostgresAdapter) DeleteList(ctx context.Context, id int64) error {
	return s.queries.DeleteList(ctx, id)
}

func (s PostgresAdapter) AllItemsFromList(ctx context.Context, listID int64, limit int32) ([]adapter.ListItem, error) {
	if items, err := s.queries.AllItemsFromList(ctx, sqlc.AllItemsFromListParams{ListID: listID, Limit: limit}); err != nil {
		return nil, fmt.Errorf("failed to get items from list: %w", err)
	} else {
		var adaptedListItems = make([]adapter.ListItem, 0, len(items))
		for _, item := range items {
			adaptedListItems = append(adaptedListItems, adapter.ListItem(item))
		}
		return adaptedListItems, nil
	}
}

func (s PostgresAdapter) CreateListItem(ctx context.Context, listID int64, content string, sort int64) (adapter.ListItem, error) {
	if item, err := s.queries.CreateListItem(ctx, sqlc.CreateListItemParams{ListID: listID, Content: content, Sort: sort}); err != nil {
		return adapter.ListItem{}, fmt.Errorf("failed to create list item: %w", err)
	} else {
		return adapter.ListItem(item), nil
	}
}

func (s PostgresAdapter) DeleteListItem(ctx context.Context, id int64) error {
	return s.queries.DeleteListItem(ctx, id)
}

func (s PostgresAdapter) GetListItem(ctx context.Context, id int64) (adapter.ListItem, error) {
	if item, err := s.queries.GetListItem(ctx, id); err != nil {
		return adapter.ListItem{}, fmt.Errorf("failed to get list item: %w", err)
	} else {
		return adapter.ListItem(item), nil
	}
}

func (s PostgresAdapter) UpdateListItemChecked(ctx context.Context, id int64, checked bool) (adapter.ListItem, error) {
	if item, err := s.queries.UpdateListItemChecked(ctx, sqlc.UpdateListItemCheckedParams{ID: id, Checked: checked}); err != nil {
		return adapter.ListItem{}, fmt.Errorf("failed to update list item checked: %w", err)
	} else {
		return adapter.ListItem(item), nil
	}
}

func (s PostgresAdapter) UpdateListItemSort(ctx context.Context, id int64, sort int64) (adapter.ListItem, error) {
	if item, err := s.queries.UpdateListItemSort(ctx, sqlc.UpdateListItemSortParams{ID: id, Sort: sort}); err != nil {
		return adapter.ListItem{}, fmt.Errorf("failed to update list item sort: %w", err)
	} else {
		return adapter.ListItem(item), nil
	}
}

func (s PostgresAdapter) UpdateListItemText(ctx context.Context, id int64, content string) (adapter.ListItem, error) {
	if item, err := s.queries.UpdateListItemText(ctx, sqlc.UpdateListItemTextParams{ID: id, Content: content}); err != nil {
		return adapter.ListItem{}, fmt.Errorf("failed to update list item text: %w", err)
	} else {
		return adapter.ListItem(item), nil
	}
}
