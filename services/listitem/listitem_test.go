package listitem_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/HomeList/db/sqlite"
	"github.com/superlinkx/HomeList/services/listitem"
	"github.com/superlinkx/HomeList/services/listitem/mocks"
)

var (
	errUnhappyPath   = errors.New("unhappy path")
	happySqlListItem = sqlite.ListItem{
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

func TestFetchListItem(t *testing.T) {
	var (
		mockQueries = mocks.NewQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().GetListItem(context.Background(), int64(1)).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.FetchListItem(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().GetListItem(context.Background(), int64(1)).
			Return(sqlite.ListItem{}, errUnhappyPath).Times(1)

		_, err := srv.FetchListItem(context.Background(), 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestFetchAllItemsFromList(t *testing.T) {
	var (
		mockQueries            = mocks.NewQueries(t)
		srv                    = listitem.NewService(mockQueries)
		allItemsFromListParams = sqlite.AllItemsFromListParams{
			ListID: 1,
			Limit:  1,
		}
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().AllItemsFromList(context.Background(), allItemsFromListParams).
			Return([]sqlite.ListItem{happySqlListItem}, nil).Times(1)

		lis, err := srv.FetchAllItemsFromList(context.Background(), 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, []listitem.ListItem{
			happyServiceListItem,
		}, lis)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().AllItemsFromList(context.Background(), allItemsFromListParams).
			Return(nil, errUnhappyPath).Times(1)

		_, err := srv.FetchAllItemsFromList(context.Background(), 1, 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestAddItemToList(t *testing.T) {
	var (
		mockQueries          = mocks.NewQueries(t)
		srv                  = listitem.NewService(mockQueries)
		createListItemParams = sqlite.CreateListItemParams{
			ListID:  1,
			Content: "test",
			Sort:    1,
		}
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().CreateListItem(context.Background(), createListItemParams).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.AddItemToList(context.Background(), 1, "test", 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().CreateListItem(context.Background(), createListItemParams).
			Return(sqlite.ListItem{}, errUnhappyPath).Times(1)

		_, err := srv.AddItemToList(context.Background(), 1, "test", 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestUpdateListItemContent(t *testing.T) {
	var (
		mockQueries              = mocks.NewQueries(t)
		srv                      = listitem.NewService(mockQueries)
		updateListItemTextParams = sqlite.UpdateListItemTextParams{
			ID:      1,
			Content: "test",
		}
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemText(context.Background(), updateListItemTextParams).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.UpdateListItemContent(context.Background(), 1, "test")
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemText(context.Background(), updateListItemTextParams).
			Return(sqlite.ListItem{}, errUnhappyPath).Times(1)

		_, err := srv.UpdateListItemContent(context.Background(), 1, "test")
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestUpdateListItemSort(t *testing.T) {
	var (
		mockQueries              = mocks.NewQueries(t)
		srv                      = listitem.NewService(mockQueries)
		updateListItemSortParams = sqlite.UpdateListItemSortParams{
			ID:   1,
			Sort: 1,
		}
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemSort(context.Background(), updateListItemSortParams).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.UpdateListItemSort(context.Background(), 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemSort(context.Background(), updateListItemSortParams).
			Return(sqlite.ListItem{}, errUnhappyPath).Times(1)

		_, err := srv.UpdateListItemSort(context.Background(), 1, 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestUpdateListItemChecked(t *testing.T) {
	var (
		mockQueries                 = mocks.NewQueries(t)
		srv                         = listitem.NewService(mockQueries)
		updateListItemCheckedParams = sqlite.UpdateListItemCheckedParams{
			ID:      1,
			Checked: false,
		}
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemChecked(context.Background(), updateListItemCheckedParams).
			Return(happySqlListItem, nil).Times(1)

		li, err := srv.UpdateListItemChecked(context.Background(), 1, false)
		assert.Nil(t, err)
		assert.Equal(t, happyServiceListItem, li)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().UpdateListItemChecked(context.Background(), updateListItemCheckedParams).
			Return(sqlite.ListItem{}, errUnhappyPath).Times(1)

		_, err := srv.UpdateListItemChecked(context.Background(), 1, false)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}

func TestDeleteListItem(t *testing.T) {
	var (
		mockQueries = mocks.NewQueries(t)
		srv         = listitem.NewService(mockQueries)
	)

	t.Run("happy path", func(t *testing.T) {
		mockQueries.EXPECT().DeleteListItem(context.Background(), int64(1)).
			Return(nil).Times(1)

		err := srv.DeleteListItem(context.Background(), 1)
		assert.Nil(t, err)
	})

	t.Run("unhappy path", func(t *testing.T) {
		mockQueries.EXPECT().DeleteListItem(context.Background(), int64(1)).
			Return(errUnhappyPath).Times(1)

		err := srv.DeleteListItem(context.Background(), 1)
		assert.ErrorIs(t, err, errUnhappyPath)
	})
}
