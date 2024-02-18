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

package app_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/app/model"
	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/data/adapter/mocks"
)

var (
	successNoOffsetItems = []model.Item{
		{ID: 0, ListID: 0, Content: "Item 0", Checked: false, Sort: 1024},
		{ID: 1, ListID: 0, Content: "Item 1", Checked: false, Sort: 2048},
		{ID: 2, ListID: 0, Content: "Item 2", Checked: false, Sort: 3072},
		{ID: 3, ListID: 0, Content: "Item 3", Checked: false, Sort: 4096},
		{ID: 4, ListID: 0, Content: "Item 4", Checked: false, Sort: 5120},
		{ID: 5, ListID: 0, Content: "Item 5", Checked: false, Sort: 6144},
		{ID: 6, ListID: 0, Content: "Item 6", Checked: false, Sort: 7168},
		{ID: 7, ListID: 0, Content: "Item 7", Checked: false, Sort: 8192},
		{ID: 8, ListID: 0, Content: "Item 8", Checked: false, Sort: 9216},
		{ID: 9, ListID: 0, Content: "Item 9", Checked: false, Sort: 10240},
	}

	successOffset5Items = []model.Item{
		{ID: 5, ListID: 0, Content: "Item 5", Checked: false, Sort: 6144},
		{ID: 6, ListID: 0, Content: "Item 6", Checked: false, Sort: 7168},
		{ID: 7, ListID: 0, Content: "Item 7", Checked: false, Sort: 8192},
		{ID: 8, ListID: 0, Content: "Item 8", Checked: false, Sort: 9216},
		{ID: 9, ListID: 0, Content: "Item 9", Checked: false, Sort: 10240},
		{ID: 10, ListID: 0, Content: "Item 10", Checked: false, Sort: 11264},
		{ID: 11, ListID: 0, Content: "Item 11", Checked: false, Sort: 12288},
		{ID: 12, ListID: 0, Content: "Item 12", Checked: false, Sort: 13312},
		{ID: 13, ListID: 0, Content: "Item 13", Checked: false, Sort: 14336},
		{ID: 14, ListID: 0, Content: "Item 14", Checked: false, Sort: 15360},
	}
)

func TestApp_AllItems(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("No items found", func(t *testing.T) {
		mockAdapter.EXPECT().AllItemsFromList(context.TODO(), int64(0), int32(10), int32(0)).Return([]model.Item{}, adapter.ErrNotFound).Times(1)
		lists, err := appInst.AllItemsFromList(context.TODO(), 0, 10, 0)
		assert.ErrorIs(t, err, app.ErrNotFound)
		assert.Empty(t, lists)
	})

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().AllItemsFromList(context.TODO(), int64(0), int32(10), int32(0)).Return([]model.Item{}, errGeneric).Times(1)
		lists, err := appInst.AllItemsFromList(context.TODO(), 0, 10, 0)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, lists)
	})

	t.Run("Generic error with nil slice", func(t *testing.T) {
		mockAdapter.EXPECT().AllItemsFromList(context.TODO(), int64(0), int32(10), int32(0)).Return(nil, errGeneric).Times(1)
		lists, err := appInst.AllItemsFromList(context.TODO(), 0, 10, 0)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, lists)
	})

	t.Run("Success with no offset but nil slice returned", func(t *testing.T) {
		mockAdapter.EXPECT().AllItemsFromList(context.TODO(), int64(0), int32(10), int32(0)).Return(nil, nil).Times(1)
		lists, err := appInst.AllItemsFromList(context.TODO(), 0, 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, lists)
	})

	t.Run("Success with no offset", func(t *testing.T) {
		mockAdapter.EXPECT().AllItemsFromList(context.TODO(), int64(0), int32(10), int32(0)).Return(successNoOffsetItems, nil).Times(1)
		lists, err := appInst.AllItemsFromList(context.TODO(), 0, 10, 0)
		assert.NoError(t, err)
		assert.Equal(t, 10, len(lists))
		for idx, list := range lists {
			assert.Equal(t, successNoOffsetLists[idx].ID, list.ID)
		}
	})

	t.Run("Success with offset 5", func(t *testing.T) {
		mockAdapter.EXPECT().AllItemsFromList(context.TODO(), int64(0), int32(10), int32(5)).Return(successOffset5Items, nil).Times(1)
		lists, err := appInst.AllItemsFromList(context.TODO(), 0, 10, 5)
		assert.NoError(t, err)
		assert.Equal(t, 10, len(lists))
		for idx, list := range lists {
			assert.Equal(t, successOffset5Lists[idx].ID, list.ID)
		}
	})
}

func TestApp_GetItem(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Item not found", func(t *testing.T) {
		mockAdapter.EXPECT().GetItemFromList(context.TODO(), int64(0), int64(0)).Return(model.Item{}, adapter.ErrNotFound).Times(1)
		list, err := appInst.GetItemFromList(context.TODO(), 0, 0)
		assert.ErrorIs(t, err, app.ErrNotFound)
		assert.Empty(t, list)
	})

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().GetItemFromList(context.TODO(), int64(0), int64(0)).Return(model.Item{}, errGeneric).Times(1)
		list, err := appInst.GetItemFromList(context.TODO(), 0, 0)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, list)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().GetItemFromList(context.TODO(), int64(0), int64(0)).Return(model.Item{ID: 0, ListID: 0, Content: "Item 0", Checked: false, Sort: 1024}, nil).Times(1)
		list, err := appInst.GetItemFromList(context.TODO(), 0, 0)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), list.ID)
	})
}

func TestApp_CreateItem(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().CreateItemOnList(context.TODO(), int64(0), "Item 0", int64(1024)).Return(model.Item{}, errGeneric).Times(1)
		item, err := appInst.CreateItemOnList(context.TODO(), 0, "Item 0", 1024)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, item)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().CreateItemOnList(context.TODO(), int64(0), "Item 0", int64(1024)).Return(model.Item{ID: 0, ListID: 0, Content: "Item 0", Checked: false, Sort: 1024}, nil).Times(1)
		item, err := appInst.CreateItemOnList(context.TODO(), 0, "Item 0", 1024)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), item.ID)
	})
}

func TestApp_UpdateItem(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Item not found", func(t *testing.T) {
		mockAdapter.EXPECT().UpdateItemFromListContent(context.TODO(), int64(0), int64(0), "Item 0").Return(model.Item{}, adapter.ErrNotFound).Times(1)
		item, err := appInst.UpdateItemFromList(context.TODO(), 0, 0, "Item 0", false, 1024)
		assert.ErrorIs(t, err, app.ErrNotFound)
		assert.Empty(t, item)
	})

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().UpdateItemFromListContent(context.TODO(), int64(0), int64(0), "Item 0").Return(model.Item{}, errGeneric).Times(1)
		item, err := appInst.UpdateItemFromList(context.TODO(), 0, 0, "Item 0", false, 1024)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, item)
	})

	t.Run("Success", func(t *testing.T) {
		var result = model.Item{ID: 0, ListID: 0, Content: "Item 20", Checked: false, Sort: 1024}
		mockAdapter.EXPECT().UpdateItemFromListContent(context.TODO(), int64(0), int64(0), "Item 20").Return(result, nil).Times(1)
		item, err := appInst.UpdateItemFromList(context.TODO(), 0, 0, "Item 20", true, 512)
		assert.NoError(t, err)
		assert.Equal(t, result, item)
	})
}

func TestApp_DeleteItem(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Item not found", func(t *testing.T) {
		mockAdapter.EXPECT().DeleteItemFromList(context.TODO(), int64(0), int64(0)).Return(adapter.ErrNotFound).Times(1)
		err := appInst.DeleteItemFromList(context.TODO(), 0, 0)
		assert.ErrorIs(t, err, app.ErrNotFound)
	})

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().DeleteItemFromList(context.TODO(), int64(0), int64(0)).Return(errGeneric).Times(1)
		err := appInst.DeleteItemFromList(context.TODO(), 0, 0)
		assert.ErrorIs(t, err, app.ErrInternal)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().DeleteItemFromList(context.TODO(), int64(0), int64(0)).Return(nil).Times(1)
		err := appInst.DeleteItemFromList(context.TODO(), 0, 0)
		assert.NoError(t, err)
	})
}
