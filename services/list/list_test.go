package list_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/superlinkx/HomeList/db/sqlite"
	"github.com/superlinkx/HomeList/harnesses/integration"
	"github.com/superlinkx/HomeList/services/list"
)

type closer interface {
	Close() error
}

func setupService() (list.Service, closer, error) {
	if db, err := integration.ConnectDatabase(); err != nil {
		return list.Service{}, db, fmt.Errorf("failed to generate db connection: %s", err)
	} else if err := db.Ping(); err != nil {
		return list.Service{}, db, fmt.Errorf("failed to ping database: %s", err)
	} else if err := integration.ResetDatabase(db); err != nil {
		return list.Service{}, db, fmt.Errorf("failed to reset database: %s", err)
	} else {
		queries := sqlite.New(db)
		srv := list.NewService(queries)
		return srv, db, nil
	}
}

func TestAllLists(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.AllLists(context.Background(), 10)
		if err != nil {
			t.Fatalf("failed to fetch all lists: %s", err)
		}
		require.GreaterOrEqual(t, len(result), 1)
	})
}

func TestGetList(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.GetList(context.Background(), 1)
		if err != nil {
			t.Fatalf("failed to get list: %s", err)
		}
		require.Equal(t, result.Name, "Main List")
		require.Equal(t, result.ID, int64(1))
	})
}

func TestUpdateList(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.UpdateList(context.Background(), 1, "updated-list-name")
		if err != nil {
			t.Fatalf("failed to update service: %s", err)
		}
		require.Equal(t, result.Name, "updated-list-name")
	})
}

func TestDeleteList(t *testing.T) {
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
