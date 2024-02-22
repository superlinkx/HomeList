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
	"github.com/stretchr/testify/mock"
	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/app/model"
	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/data/adapter/mocks"
)

var (
	successNoOffsetLists = []model.List{
		{ID: 0, Name: "List 0"},
		{ID: 1, Name: "List 1"},
		{ID: 2, Name: "List 2"},
		{ID: 3, Name: "List 3"},
		{ID: 4, Name: "List 4"},
		{ID: 5, Name: "List 5"},
		{ID: 6, Name: "List 6"},
		{ID: 7, Name: "List 7"},
		{ID: 8, Name: "List 8"},
		{ID: 9, Name: "List 9"},
	}

	successOffset5Lists = []model.List{
		{ID: 5, Name: "List 5"},
		{ID: 6, Name: "List 6"},
		{ID: 7, Name: "List 7"},
		{ID: 8, Name: "List 8"},
		{ID: 9, Name: "List 9"},
		{ID: 10, Name: "List 10"},
		{ID: 11, Name: "List 11"},
		{ID: 12, Name: "List 12"},
		{ID: 13, Name: "List 13"},
		{ID: 14, Name: "List 14"},
	}
)

func TestApp_AllLists(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("No lists found", func(t *testing.T) {
		mockAdapter.EXPECT().AllLists(context.TODO(), int32(10), int32(0)).Return([]model.List{}, adapter.ErrNotFound).Times(1)
		lists, err := appInst.AllLists(context.TODO(), 10, 0)
		assert.ErrorIs(t, err, app.ErrNotFound)
		assert.Empty(t, lists)
	})

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().AllLists(context.TODO(), int32(10), int32(0)).Return([]model.List{}, errGeneric).Times(1)
		lists, err := appInst.AllLists(context.TODO(), 10, 0)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, lists)
	})

	t.Run("Generic error with nil slice", func(t *testing.T) {
		mockAdapter.EXPECT().AllLists(context.TODO(), int32(10), int32(0)).Return(nil, errGeneric).Times(1)
		lists, err := appInst.AllLists(context.TODO(), 10, 0)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, lists)
	})

	t.Run("Success with no offset but nil slice returned", func(t *testing.T) {
		mockAdapter.EXPECT().AllLists(context.TODO(), int32(10), int32(0)).Return(nil, nil).Times(1)
		lists, err := appInst.AllLists(context.TODO(), 10, 0)
		assert.NoError(t, err)
		assert.Empty(t, lists)
	})

	t.Run("Success with no offset", func(t *testing.T) {
		mockAdapter.EXPECT().AllLists(context.TODO(), int32(10), int32(0)).Return(successNoOffsetLists, nil).Times(1)
		lists, err := appInst.AllLists(context.TODO(), 10, 0)
		assert.NoError(t, err)
		assert.Equal(t, 10, len(lists))
		for idx, list := range lists {
			assert.Equal(t, successNoOffsetLists[idx].ID, list.ID)
		}
	})

	t.Run("Success with offset 5", func(t *testing.T) {
		mockAdapter.EXPECT().AllLists(context.TODO(), int32(10), int32(5)).Return(successOffset5Lists, nil).Times(1)
		lists, err := appInst.AllLists(context.TODO(), 10, 5)
		assert.NoError(t, err)
		assert.Equal(t, 10, len(lists))
		for idx, list := range lists {
			assert.Equal(t, successOffset5Lists[idx].ID, list.ID)
		}
	})
}

func TestApp_GetList(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("List not found", func(t *testing.T) {
		mockAdapter.EXPECT().GetList(context.TODO(), int64(1)).Return(model.List{}, adapter.ErrNotFound).Times(1)
		list, err := appInst.GetList(context.TODO(), 1)
		assert.ErrorIs(t, err, app.ErrNotFound)
		assert.Empty(t, list)
	})

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().GetList(context.TODO(), int64(1)).Return(model.List{}, errGeneric).Times(1)
		list, err := appInst.GetList(context.TODO(), 1)
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, list)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().GetList(context.TODO(), int64(1)).Return(successNoOffsetLists[1], nil).Times(1)
		list, err := appInst.GetList(context.TODO(), 1)
		assert.NoError(t, err)
		assert.Equal(t, successNoOffsetLists[1].ID, list.ID)
	})
}

func TestApp_CreateList(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().CreateList(context.TODO(), "List 1").Return(model.List{}, errGeneric).Times(1)
		list, err := appInst.CreateList(context.TODO(), "List 1")
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, list)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().CreateList(context.TODO(), "List 1").Return(successNoOffsetLists[1], nil).Times(1)
		list, err := appInst.CreateList(context.TODO(), "List 1")
		assert.NoError(t, err)
		assert.Equal(t, successNoOffsetLists[1].ID, list.ID)
	})
}

func TestApp_UpdateList(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().RenameList(context.TODO(), int64(1), "List 1").Return(model.List{}, errGeneric).Times(1)
		list, err := appInst.UpdateList(context.TODO(), 1, "List 1")
		assert.ErrorIs(t, err, app.ErrInternal)
		assert.Empty(t, list)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().RenameList(context.TODO(), int64(1), "List 1").Return(successNoOffsetLists[1], nil).Times(1)
		list, err := appInst.UpdateList(context.TODO(), 1, "List 1")
		assert.NoError(t, err)
		assert.Equal(t, successNoOffsetLists[1].ID, list.ID)
	})
}

func TestApp_DeleteList(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().DeleteList(context.TODO(), int64(1)).Return(errGeneric).Times(1)
		err := appInst.DeleteList(context.TODO(), 1)
		assert.ErrorIs(t, err, app.ErrInternal)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().DeleteList(context.TODO(), int64(1)).Return(nil).Times(1)
		err := appInst.DeleteList(context.TODO(), 1)
		assert.NoError(t, err)
	})
}

func TestApp_ReflowList(t *testing.T) {
	var (
		mockAdapter = mocks.NewMockAdapter(t)
		appInst     = app.NewApp(mockAdapter)
	)

	t.Parallel()

	t.Run("Generic error", func(t *testing.T) {
		mockAdapter.EXPECT().ReflowList(context.TODO(), int64(1), mock.AnythingOfType("mapper.Reflow")).Return(errGeneric).Times(1)
		err := appInst.ReflowList(context.TODO(), 1)
		assert.ErrorIs(t, err, app.ErrInternal)
	})

	t.Run("Success", func(t *testing.T) {
		mockAdapter.EXPECT().ReflowList(context.TODO(), int64(1), mock.AnythingOfType("mapper.Reflow")).Return(nil).Times(1)
		err := appInst.ReflowList(context.TODO(), 1)
		assert.NoError(t, err)
	})
}
