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
	"github.com/superlinkx/HomeList/app/listitem"
	"github.com/superlinkx/HomeList/app/listitem/mocks"
	listitemsrv "github.com/superlinkx/HomeList/service/listitem"
)

var (
	errServiceFailure   = errors.New("service failure")
	listItemSrvListItem = listitemsrv.ListItem{
		ID:      1,
		ListID:  1,
		Content: "Test",
		Checked: true,
		Sort:    1,
	}
	listItemSrvListItemSlice = []listitemsrv.ListItem{listItemSrvListItem}
	listItemAppListItem      = listitem.ListItem(listItemSrvListItem)
	listItemAppListItemSlice = []listitem.ListItem{listItemAppListItem}
)

func TestListItemApp_FetchListItem(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().FetchListItem(context.Background(), int64(1)).
			Return(listItemSrvListItem, nil).Times(1)

		item, err := listItemApp.FetchListItem(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, listItemAppListItem, item)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().FetchListItem(context.Background(), int64(1)).
			Return(listitemsrv.ListItem{}, errServiceFailure).Times(1)

		_, err := listItemApp.FetchListItem(context.Background(), 1)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListItemApp_FetchAllItemsFromList(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().FetchAllItemsFromList(context.Background(), int64(1), int32(10)).
			Return(listItemSrvListItemSlice, nil).Times(1)

		items, err := listItemApp.FetchAllItemsFromList(context.Background(), 1, 10)
		assert.Nil(t, err)
		assert.Equal(t, listItemAppListItemSlice, items)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().FetchAllItemsFromList(context.Background(), int64(1), int32(10)).
			Return(nil, errServiceFailure).Times(1)

		_, err := listItemApp.FetchAllItemsFromList(context.Background(), 1, 10)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListItemApp_AddItemToList(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().AddItemToList(context.Background(), listItemAppListItem.ListID, listItemAppListItem.Content, listItemAppListItem.Sort).
			Return(listItemSrvListItem, nil).Times(1)

		item, err := listItemApp.AddItemToList(context.Background(), listItemAppListItem.ListID, listItemAppListItem.Content, listItemAppListItem.Sort)
		assert.Nil(t, err)
		assert.Equal(t, listItemAppListItem, item)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().AddItemToList(context.Background(), listItemAppListItem.ListID, listItemAppListItem.Content, listItemAppListItem.Sort).
			Return(listitemsrv.ListItem{}, errServiceFailure).Times(1)

		_, err := listItemApp.AddItemToList(context.Background(), listItemAppListItem.ListID, listItemAppListItem.Content, listItemAppListItem.Sort)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListItemApp_UpdateListItemContent(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().UpdateListItemContent(context.Background(), listItemAppListItem.ID, listItemAppListItem.Content).
			Return(listItemSrvListItem, nil).Times(1)

		item, err := listItemApp.UpdateListItemContent(context.Background(), listItemAppListItem.ID, listItemAppListItem.Content)
		assert.Nil(t, err)
		assert.Equal(t, listItemAppListItem, item)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().UpdateListItemContent(context.Background(), listItemAppListItem.ID, listItemAppListItem.Content).
			Return(listitemsrv.ListItem{}, errServiceFailure).Times(1)

		_, err := listItemApp.UpdateListItemContent(context.Background(), listItemAppListItem.ID, listItemAppListItem.Content)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListItemApp_UpdateListItemChecked(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().UpdateListItemChecked(context.Background(), listItemAppListItem.ID, listItemAppListItem.Checked).
			Return(listItemSrvListItem, nil).Times(1)

		item, err := listItemApp.UpdateListItemChecked(context.Background(), listItemAppListItem.ID, listItemAppListItem.Checked)
		assert.Nil(t, err)
		assert.Equal(t, listItemAppListItem, item)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().UpdateListItemChecked(context.Background(), listItemAppListItem.ID, listItemAppListItem.Checked).
			Return(listitemsrv.ListItem{}, errServiceFailure).Times(1)

		_, err := listItemApp.UpdateListItemChecked(context.Background(), listItemAppListItem.ID, listItemAppListItem.Checked)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListItemApp_UpdateListItemSort(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().UpdateListItemSort(context.Background(), listItemAppListItem.ID, listItemAppListItem.Sort).
			Return(listItemSrvListItem, nil).Times(1)

		item, err := listItemApp.UpdateListItemSort(context.Background(), listItemAppListItem.ID, listItemAppListItem.Sort)
		assert.Nil(t, err)
		assert.Equal(t, listItemAppListItem, item)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().UpdateListItemSort(context.Background(), listItemAppListItem.ID, listItemAppListItem.Sort).
			Return(listitemsrv.ListItem{}, errServiceFailure).Times(1)

		_, err := listItemApp.UpdateListItemSort(context.Background(), listItemAppListItem.ID, listItemAppListItem.Sort)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListItemApp_DeleteListItem(t *testing.T) {
	var (
		mockListItemService = mocks.NewMockListItemService(t)
		listItemApp         = listitem.NewListItemApp(mockListItemService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListItemService.EXPECT().DeleteListItem(context.Background(), listItemAppListItem.ID).
			Return(nil).Times(1)

		err := listItemApp.DeleteListItem(context.Background(), listItemAppListItem.ID)
		assert.Nil(t, err)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListItemService.EXPECT().DeleteListItem(context.Background(), listItemAppListItem.ID).
			Return(errServiceFailure).Times(1)

		err := listItemApp.DeleteListItem(context.Background(), listItemAppListItem.ID)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}
