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

	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/data/sqlite/sqlc"
	"github.com/superlinkx/HomeList/model"
)

type SqliteAdapter struct {
	queries *sqlc.Queries
}

func NewSqliteAdapter(db *sql.DB) adapter.Adapter {
	return SqliteAdapter{queries: sqlc.New(db)}
}

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
	return s.queries.DeleteList(ctx, id)
}
