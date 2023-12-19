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

package handler

import (
	"context"

	"github.com/superlinkx/HomeList/app"
)

type Handlers struct {
	ListHandlers
	ListItemHandlers
}

type ListHandlers struct {
	list list
}

type ListItemHandlers struct {
	listItem listItem
}

type application interface {
	list
	listItem
}

type list interface {
	AllLists(context.Context, int64) ([]app.List, error)
	GetList(context.Context, int64) (app.List, error)
	CreateList(context.Context, string) (app.List, error)
	UpdateList(context.Context, int64, string) (app.List, error)
	DeleteList(context.Context, int64) error
}

type listItem interface {
	FetchAllItemsFromList(context.Context, int64, int64) ([]app.ListItem, error)
	FetchListItem(context.Context, int64) (app.ListItem, error)
	AddItemToList(context.Context, int64, string, int64) (app.ListItem, error)
	UpdateListItemContent(context.Context, int64, string) (app.ListItem, error)
	UpdateListItemChecked(context.Context, int64, bool) (app.ListItem, error)
	UpdateListItemSort(context.Context, int64, int64) (app.ListItem, error)
	DeleteListItem(context.Context, int64) error
}

func NewHandlersWithApplication(a application) Handlers {
	return Handlers{
		ListHandlers: ListHandlers{
			list: a,
		},
		ListItemHandlers: ListItemHandlers{
			listItem: a,
		},
	}
}

func NewHandlersWithComponents(l list, li listItem) Handlers {
	return Handlers{
		ListHandlers: ListHandlers{
			list: l,
		},
		ListItemHandlers: ListItemHandlers{
			listItem: li,
		},
	}
}
