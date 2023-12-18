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
	"github.com/superlinkx/HomeList/db/sqlite"
	"github.com/superlinkx/HomeList/services/list"
	"github.com/superlinkx/HomeList/services/list/mocks"
)

var (
	errUnhappyPath = errors.New("unhappy path")
	happySqlList   = sqlite.List{
		ID:   1,
		Name: "test",
	}
	happyServiceList = list.List{
		ID:   1,
		Name: "test",
	}
)

func TestGetList(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = list.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().GetList(context.Background(), int64(1)).
			Return(happySqlList, nil).Times(1)

		l, err := srv.GetList(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceList, l)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().GetList(context.Background(), int64(1)).
			Return(sqlite.List{}, errUnhappyPath).Times(1)

		_, err := srv.GetList(context.Background(), 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestAllLists(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = list.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().AllLists(context.Background(), int64(1)).
			Return([]sqlite.List{happySqlList}, nil).Times(1)

		l, err := srv.AllLists(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, []list.List{
			happyServiceList,
		}, l)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().AllLists(context.Background(), int64(1)).
			Return(nil, errUnhappyPath).Times(1)

		_, err := srv.AllLists(context.Background(), 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestCreateList(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = list.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().CreateList(context.Background(), "test").
			Return(happySqlList, nil).Times(1)

		l, err := srv.CreateList(context.Background(), "test")
		assert.Nil(t, err)
		assert.Equal(t, happyServiceList, l)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().CreateList(context.Background(), "test").
			Return(sqlite.List{}, errUnhappyPath).Times(1)

		_, err := srv.CreateList(context.Background(), "test")
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestUpdateList(t *testing.T) {
	var (
		mockQueries      = mocks.NewMockQueries(t)
		srv              = list.NewService(mockQueries)
		renameListParams = sqlite.RenameListParams{
			ID:   1,
			Name: "test",
		}
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().RenameList(context.Background(), renameListParams).
			Return(happySqlList, nil).Times(1)

		l, err := srv.UpdateList(context.Background(), 1, "test")
		assert.Nil(t, err)
		assert.Equal(t, happyServiceList, l)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().RenameList(context.Background(), renameListParams).
			Return(sqlite.List{}, errUnhappyPath).Times(1)

		_, err := srv.UpdateList(context.Background(), 1, "test")
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestDeleteList(t *testing.T) {
	var (
		mockQueries = mocks.NewMockQueries(t)
		srv         = list.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().DeleteList(context.Background(), int64(1)).
			Return(nil).Times(1)

		err := srv.DeleteList(context.Background(), 1)
		assert.Nil(t, err)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().DeleteList(context.Background(), int64(1)).
			Return(errUnhappyPath).Times(1)

		err := srv.DeleteList(context.Background(), 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}
