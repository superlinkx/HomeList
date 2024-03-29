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

package app

import (
	"context"
	"errors"
	"log/slog"

	"github.com/superlinkx/HomeList/app/model"
	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/pkg/mapper"
)

func (s App) AllLists(ctx context.Context, limit int32, offset int32) ([]model.List, error) {
	if lists, err := s.adapter.AllLists(ctx, limit, offset); errors.Is(err, adapter.ErrNotFound) {
		slog.Error("No lists found", "error", err)
		return nil, ErrNotFound
	} else if err != nil {
		slog.Error("Unexpected error getting lists", "error", err)
		return nil, ErrInternal
	} else {
		return lists, nil
	}
}

func (s App) GetList(ctx context.Context, id int64) (model.List, error) {
	if list, err := s.adapter.GetList(ctx, id); errors.Is(err, adapter.ErrNotFound) {
		slog.Error("No list found", "id", id, "error", err)
		return model.List{}, ErrNotFound
	} else if err != nil {
		slog.Error("Unexpected error getting list", "id", id, "error", err)
		return model.List{}, ErrInternal
	} else {
		return list, nil
	}
}

func (s App) CreateList(ctx context.Context, name string) (model.List, error) {
	if list, err := s.adapter.CreateList(ctx, name); err != nil {
		slog.Error("Unexpected error creating list", "name", name, "error", err)
		return model.List{}, ErrInternal
	} else {
		return list, nil
	}
}

func (s App) UpdateList(ctx context.Context, id int64, name string) (model.List, error) {
	if list, err := s.adapter.RenameList(ctx, id, name); errors.Is(err, adapter.ErrNotFound) {
		slog.Error("No list found", "id", id, "error", err)
		return model.List{}, ErrNotFound
	} else if err != nil {
		slog.Error("Unexpected error renaming list", "id", id, "name", name, "error", err)
		return model.List{}, ErrInternal
	} else {
		return list, nil
	}
}

func (s App) DeleteList(ctx context.Context, id int64) error {
	if err := s.adapter.DeleteList(ctx, id); errors.Is(err, adapter.ErrNotFound) {
		slog.Error("No list found", "id", id, "error", err)
		return ErrNotFound
	} else if err != nil {
		slog.Error("Unexpected error deleting list", "id", id, "error", err)
		return ErrInternal
	} else {
		return nil
	}
}

func (s App) ReflowList(ctx context.Context, id int64) error {
	if err := s.adapter.ReflowList(ctx, id, mapper.ReflowListOfItemsWithSpacing(1024)); err != nil {
		slog.Error("Unexpected error reflowing list", "id", id, "error", err)
		return ErrInternal
	} else {
		return nil
	}
}
