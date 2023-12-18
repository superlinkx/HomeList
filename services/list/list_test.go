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
		mockQueries = mocks.NewQueries(t)
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
		mockQueries = mocks.NewQueries(t)
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
		mockQueries = mocks.NewQueries(t)
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
		mockQueries      = mocks.NewQueries(t)
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
		mockQueries = mocks.NewQueries(t)
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
