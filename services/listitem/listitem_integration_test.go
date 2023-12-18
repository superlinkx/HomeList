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
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/superlinkx/HomeList/db/sqlite"
	"github.com/superlinkx/HomeList/harnesses/integration"
	"github.com/superlinkx/HomeList/services/listitem"
)

type closer interface {
	Close() error
}

func setupService() (listitem.Service, closer, error) {
	if db, err := integration.ConnectDatabase(); err != nil {
		return listitem.Service{}, db, fmt.Errorf("failed to generate db connection: %s", err)
	} else if err := db.Ping(); err != nil {
		return listitem.Service{}, db, fmt.Errorf("failed to ping database: %s", err)
	} else if err := integration.ResetDatabase(db); err != nil {
		return listitem.Service{}, db, fmt.Errorf("failed to reset database: %s", err)
	} else {
		queries := sqlite.New(db)
		srv := listitem.NewService(queries)
		return srv, db, nil
	}
}

func TestIntegration_FetchListItem(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		li, err := srv.FetchListItem(context.Background(), 1)
		require.Nilf(t, err, "failed to fetch list item: %s", err)

		require.Equal(t, int64(1), li.ID)
		require.Equal(t, int64(1), li.ListID)
		require.Equal(t, int64(1024), li.Sort)
		require.Equal(t, "Hello world!", li.Content)
		require.Equal(t, false, li.Checked)
	})
}

func TestIntegration_FetchAllItemsFromList(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.FetchAllItemsFromList(context.Background(), 1, 10)
		require.Nilf(t, err, "failed to fetch list items: %s", err)
		require.GreaterOrEqual(t, len(result), 1)

		li := result[0]
		require.Equal(t, int64(1), li.ID)
		require.Equal(t, int64(1), li.ListID)
		require.Equal(t, int64(1024), li.Sort)
		require.Equal(t, "Hello world!", li.Content)
		require.Equal(t, false, li.Checked)
	})

	t.Run("NonExistentList", func(t *testing.T) {
		result, err := srv.FetchAllItemsFromList(context.Background(), 0, 10)
		require.Nilf(t, err, "failed to fetch all items: %s", err)
		require.Len(t, result, 0)
	})
}

func TestIntegration_AddItemToList(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.AddItemToList(context.Background(), 1, "test-content", 1)
		require.Nilf(t, err, "failed to add item to list: %s", err)
		require.Equal(t, result.Content, "test-content")
		require.Equal(t, result.Sort, int64(1))
		require.Equal(t, result.ListID, int64(1))

		listResult, err := srv.FetchAllItemsFromList(context.Background(), 1, 10)
		require.Nilf(t, err, "failed to fetch list items: %s", err)
		require.GreaterOrEqual(t, len(listResult), 2)
		require.Contains(t, listResult, result)
	})
}

func TestIntegration_UpdateListItem(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service")
	defer closer.Close()

	t.Run("ContentHappyPath", func(t *testing.T) {
		result, err := srv.UpdateListItemContent(context.Background(), 1, "updated-content")
		require.Nilf(t, err, "failed to update list item content: %s", err)
		require.Equal(t, result.Content, "updated-content")
	})

	t.Run("SortHappyPath", func(t *testing.T) {
		result, err := srv.UpdateListItemSort(context.Background(), 1, 2)
		require.Nilf(t, err, "failed to update list item sort: %s", err)
		require.Equal(t, result.Sort, int64(2))
	})

	t.Run("Checked HappyPath,", func(t *testing.T) {
		result, err := srv.UpdateListItemChecked(context.Background(), 1, true)
		require.Nilf(t, err, "failed to update list item checked: %s", err)
		require.Equal(t, result.Checked, true)
	})
}

func TestIntegration_DeleteListItem(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		itemResult, err := srv.FetchListItem(context.Background(), 1)
		require.Nilf(t, err, "failed to fetch list item: %s", err)

		err = srv.DeleteListItem(context.Background(), itemResult.ID)
		require.Nilf(t, err, "failed to delete list item: %s", err)

		listResult, err := srv.FetchAllItemsFromList(context.Background(), 1, 10)
		require.Nilf(t, err, "failed to fetch list items: %s", err)
		require.NotContains(t, listResult, itemResult)
	})
}
