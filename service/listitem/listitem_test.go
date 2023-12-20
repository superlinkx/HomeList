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

package listitem_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/service/listitem"
	"github.com/superlinkx/HomeList/service/listitem/mocks"
)

var (
	errAdapterFailure = errors.New("adapter failure")
	happySqlListItem  = adapter.ListItem{
		ID:      1,
		ListID:  1,
		Content: "test",
		Sort:    1,
		Checked: false,
	}
	happyServiceListItem = listitem.ListItem{
		ID:      1,
		ListID:  1,
		Content: "test",
		Sort:    1,
		Checked: false,
	}
)

func TestListItemService_FetchListItem(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().GetListItem(context.Background(), int64(1)).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.FetchListItem(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().GetListItem(context.Background(), int64(1)).
			Return(adapter.ListItem{}, errAdapterFailure).Times(1)

		_, err := srv.FetchListItem(context.Background(), 1)
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}

func TestListItemService_FetchAllItemsFromList(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().AllItemsFromList(context.Background(), int64(1), int32(1)).
			Return([]adapter.ListItem{happySqlListItem}, nil).Times(1)

		lis, err := srv.FetchAllItemsFromList(context.Background(), 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, []listitem.ListItem{
			happyServiceListItem,
		}, lis)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().AllItemsFromList(context.Background(), int64(1), int32(1)).
			Return(nil, errAdapterFailure).Times(1)

		_, err := srv.FetchAllItemsFromList(context.Background(), 1, 1)
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}

func TestListItemService_AddItemToList(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().CreateListItem(context.Background(), int64(1), "test", int64(1)).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.AddItemToList(context.Background(), 1, "test", 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().CreateListItem(context.Background(), int64(1), "test", int64(1)).
			Return(adapter.ListItem{}, errAdapterFailure).Times(1)

		_, err := srv.AddItemToList(context.Background(), 1, "test", 1)
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}

func TestListItemService_UpdateListItemContent(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemText(context.Background(), int64(1), "test").
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.UpdateListItemContent(context.Background(), 1, "test")
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemText(context.Background(), int64(1), "test").
			Return(adapter.ListItem{}, errAdapterFailure).Times(1)

		_, err := srv.UpdateListItemContent(context.Background(), 1, "test")
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}

func TestListItemService_UpdateListItemSort(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemSort(context.Background(), int64(1), int64(1)).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.UpdateListItemSort(context.Background(), 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemSort(context.Background(), int64(1), int64(1)).
			Return(adapter.ListItem{}, errAdapterFailure).Times(1)

		_, err := srv.UpdateListItemSort(context.Background(), 1, 1)
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}

func TestListItemService_UpdateListItemChecked(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemChecked(context.Background(), int64(1), false).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.UpdateListItemChecked(context.Background(), 1, false)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemChecked(context.Background(), int64(1), false).
			Return(adapter.ListItem{}, errAdapterFailure).Times(1)

		_, err := srv.UpdateListItemChecked(context.Background(), 1, false)
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}

func TestListItemService_DeleteListItem(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().DeleteListItem(context.Background(), int64(1)).
			Return(nil).Times(1)

		err := srv.DeleteListItem(context.Background(), 1)
		assert.Nil(t, err)
	})

	t.Run("adapter failure", func(t *testing.T) {
		mockQueries.EXPECT().DeleteListItem(context.Background(), int64(1)).
			Return(errAdapterFailure).Times(1)

		err := srv.DeleteListItem(context.Background(), 1)
		assert.ErrorIs(t, err, errAdapterFailure)
	})
}
