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
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/superlinkx/HomeList/data/sqlite/adapter"
	"github.com/superlinkx/HomeList/harnesses/integration"
	"github.com/superlinkx/HomeList/service/list"
)

type closer interface {
	Close() error
}

func setupService() (list.ListService, closer, error) {
	if db, err := integration.ConnectDatabase(); err != nil {
		return list.ListService{}, db, fmt.Errorf("failed to generate db connection: %s", err)
	} else if err := db.Ping(); err != nil {
		return list.ListService{}, db, fmt.Errorf("failed to ping database: %s", err)
	} else if err := integration.ResetDatabase(db); err != nil {
		return list.ListService{}, db, fmt.Errorf("failed to reset database: %s", err)
	} else {
		queries := adapter.NewSqliteAdapter(db)
		srv := list.NewService(queries)
		return srv, db, nil
	}
}

func TestIntegration_AllLists(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.AllLists(context.Background(), 10)
		require.Nilf(t, err, "failed to fetch all lists: %s", err)
		require.GreaterOrEqual(t, len(result), 1)
	})
}

func TestIntegration_GetList(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.GetList(context.Background(), 1)
		require.Nilf(t, err, "failed to fetch list: %s", err)
		require.Equal(t, result.Name, "Main List")
		require.Equal(t, result.ID, int64(1))
	})
}

func TestIntegration_UpdateList(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.UpdateList(context.Background(), 1, "updated-list-name")
		require.Nilf(t, err, "failed to update list: %s", err)
		require.Equal(t, result.Name, "updated-list-name")
	})
}

func TestIntegration_DeleteList(t *testing.T) {
	srv, closer, err := setupService()
	require.Nilf(t, err, "failed to setup service: %s", err)
	defer closer.Close()

	t.Run("HappyPath", func(t *testing.T) {
		listResult, err := srv.GetList(context.Background(), 1)
		require.Nilf(t, err, "failed to get list")

		deleteErr := srv.DeleteList(context.Background(), 1)
		require.Nilf(t, deleteErr, "failed to delete list")

		allListsResult, err := srv.AllLists(context.Background(), 10)
		require.Nilf(t, err, "failed to get ALL lists: %s", err)
		require.NotContains(t, allListsResult, listResult)
	})
}
