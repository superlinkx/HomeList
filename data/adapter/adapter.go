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
	"errors"
)

type Adapter interface {
	ListAdapter
	ListItemAdapter
}

type ListAdapter interface {
	AllLists(ctx context.Context, limit int32) ([]List, error)
	GetList(ctx context.Context, id int64) (List, error)
	RenameList(ctx context.Context, id int64, name string) (List, error)
	CreateList(ctx context.Context, name string) (List, error)
	DeleteList(ctx context.Context, id int64) error
}

type ListItemAdapter interface {
	AllItemsFromList(ctx context.Context, listID int64, limit int32) ([]ListItem, error)
	CreateListItem(ctx context.Context, listID int64, content string, sort int64) (ListItem, error)
	DeleteListItem(ctx context.Context, id int64) error
	GetListItem(ctx context.Context, id int64) (ListItem, error)
	UpdateListItemChecked(ctx context.Context, id int64, checked bool) (ListItem, error)
	UpdateListItemSort(ctx context.Context, id int64, sort int64) (ListItem, error)
	UpdateListItemText(ctx context.Context, id int64, content string) (ListItem, error)
}

type List struct {
	ID   int64
	Name string
}

type ListItem struct {
	ID      int64
	ListID  int64
	Content string
	Checked bool
	Sort    int64
}

var (
	ErrNotFound = errors.New("not found")
)
