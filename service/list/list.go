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

package list

import (
	"context"
	"fmt"

	"github.com/superlinkx/HomeList/data/adapter"
)

type queries interface {
	AllLists(ctx context.Context, limit int32) ([]adapter.List, error)
	GetList(ctx context.Context, id int64) (adapter.List, error)
	CreateList(ctx context.Context, name string) (adapter.List, error)
	RenameList(ctx context.Context, id int64, name string) (adapter.List, error)
	DeleteList(ctx context.Context, id int64) error
}

type ListService struct {
	queries queries
}

type List struct {
	ID   int64
	Name string
}

func NewService(q queries) ListService {
	return ListService{queries: q}
}

func (s ListService) AllLists(ctx context.Context, limit int32) ([]List, error) {
	if results, err := s.queries.AllLists(ctx, limit); err != nil {
		return []List{}, fmt.Errorf("unable to get all lists: %w", err)
	} else {
		var lists = make([]List, 0, len(results))
		for _, result := range results {
			lists = append(lists, List(result))
		}
		return lists, nil
	}
}

func (s ListService) GetList(ctx context.Context, id int64) (List, error) {
	if result, err := s.queries.GetList(ctx, id); err != nil {
		return List{}, fmt.Errorf("unable to get list: %w", err)
	} else {
		return List(result), nil
	}
}

func (s ListService) CreateList(ctx context.Context, name string) (List, error) {
	if result, err := s.queries.CreateList(ctx, name); err != nil {
		return List{}, fmt.Errorf("unable to create list: %w", err)
	} else {
		return List(result), nil
	}
}

func (s ListService) UpdateList(ctx context.Context, id int64, name string) (List, error) {
	if result, err := s.queries.RenameList(ctx, id, name); err != nil {
		return List{}, fmt.Errorf("unable to update list: %w", err)
	} else {
		return List(result), nil
	}
}

func (s ListService) DeleteList(ctx context.Context, id int64) error {
	if err := s.queries.DeleteList(ctx, id); err != nil {
		return fmt.Errorf("unable to delete list: %w", err)
	} else {
		return nil
	}
}
