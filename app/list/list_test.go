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

package list_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/HomeList/app/list"
	"github.com/superlinkx/HomeList/app/list/mocks"
	listsrv "github.com/superlinkx/HomeList/service/list"
)

var (
	errServiceFailure = errors.New("service failure")
	listSrvList       = listsrv.List{
		ID:   1,
		Name: "test",
	}
	listSrvListSlice = []listsrv.List{listSrvList}
	listAppList      = list.List(listSrvList)
	listAppListSlice = []list.List{listAppList}
)

func TestListApp_AllLists(t *testing.T) {
	var (
		mockListService = mocks.NewMockListService(t)
		listApp         = list.NewListApp(mockListService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListService.EXPECT().AllLists(context.Background(), int32(10)).
			Return(listSrvListSlice, nil).Times(1)

		lists, err := listApp.AllLists(context.Background(), 10)
		assert.Nil(t, err)
		assert.Equal(t, listAppListSlice, lists)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListService.EXPECT().AllLists(context.Background(), int32(10)).
			Return(nil, errServiceFailure).Times(1)

		_, err := listApp.AllLists(context.Background(), 10)

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListApp_GetList(t *testing.T) {
	var (
		mockListService = mocks.NewMockListService(t)
		listApp         = list.NewListApp(mockListService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListService.EXPECT().GetList(context.Background(), int64(1)).
			Return(listSrvList, nil).Times(1)

		list, err := listApp.GetList(context.Background(), int64(1))
		assert.Nil(t, err)
		assert.Equal(t, listAppList, list)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListService.EXPECT().GetList(context.Background(), int64(1)).
			Return(listsrv.List{}, errServiceFailure).Times(1)

		_, err := listApp.GetList(context.Background(), int64(1))

		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListApp_CreateList(t *testing.T) {
	var (
		mockListService = mocks.NewMockListService(t)
		listApp         = list.NewListApp(mockListService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListService.EXPECT().CreateList(context.Background(), "test").
			Return(listSrvList, nil).Times(1)

		list, err := listApp.CreateList(context.Background(), "test")
		assert.Nil(t, err)
		assert.Equal(t, listAppList, list)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListService.EXPECT().CreateList(context.Background(), "test").
			Return(listsrv.List{}, errServiceFailure).Times(1)

		_, err := listApp.CreateList(context.Background(), "test")
		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListApp_UpdateList(t *testing.T) {
	var (
		mockListService = mocks.NewMockListService(t)
		listApp         = list.NewListApp(mockListService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListService.EXPECT().UpdateList(context.Background(), int64(1), "test").
			Return(listSrvList, nil).Times(1)

		list, err := listApp.UpdateList(context.Background(), int64(1), "test")
		assert.Nil(t, err)
		assert.Equal(t, listAppList, list)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListService.EXPECT().UpdateList(context.Background(), int64(1), "test").
			Return(listsrv.List{}, errServiceFailure).Times(1)

		_, err := listApp.UpdateList(context.Background(), int64(1), "test")
		assert.ErrorIs(t, err, errServiceFailure)
	})
}

func TestListApp_DeleteList(t *testing.T) {
	var (
		mockListService = mocks.NewMockListService(t)
		listApp         = list.NewListApp(mockListService)
	)

	t.Run("happy path", func(t *testing.T) {
		mockListService.EXPECT().DeleteList(context.Background(), int64(1)).
			Return(nil).Times(1)

		err := listApp.DeleteList(context.Background(), int64(1))
		assert.Nil(t, err)
	})

	t.Run("service failure", func(t *testing.T) {
		mockListService.EXPECT().DeleteList(context.Background(), int64(1)).
			Return(errServiceFailure).Times(1)

		err := listApp.DeleteList(context.Background(), int64(1))
		assert.ErrorIs(t, err, errServiceFailure)
	})
}
