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

func TestFetchListItem(t *testing.T) {
	if srv, closer, err := setupService(); err != nil {
		t.Fatalf("failed to setup service: %s", err)
	} else if li, err := srv.FetchListItem(context.Background(), 1); err != nil {
		defer closer.Close()
		t.Fatalf("failed to fetch list item: %s", err)
	} else {
		defer closer.Close()
		require.Equal(t, int64(1), li.ID)
		require.Equal(t, int64(1), li.ListID)
		require.Equal(t, int64(1024), li.Sort)
		require.Equal(t, "Hello world!", li.Content)
		require.Equal(t, false, li.Checked)
	}

}

func TestFetchAllItemsFromList(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.FetchAllItemsFromList(context.Background(), 1, 10)
		if err != nil {
			t.Fatalf("failed to fetch all items: %s", err)
		}
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
		if err != nil {
			t.Fatalf("failed to fetch all items: %s", err)
		}
		require.Len(t, result, 0)
	})
}

func TestAddItemToList(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		result, err := srv.AddItemToList(context.Background(), 1, "test-content", 1)
		if err != nil {
			t.Fatalf("failed to add item to list: %s", err)
		}
		require.Equal(t, result.Content, "test-content")
		require.Equal(t, result.Sort, int64(1))
		require.Equal(t, result.ListID, int64(1))
		listResult, err := srv.FetchAllItemsFromList(context.Background(), 1, 10)
		if err != nil {
			t.Fatalf("failed to fetch items from list: %s", err)
		}
		require.GreaterOrEqual(t, len(listResult), 2)
		require.Contains(t, listResult, result)

	})
}

func TestUpdateListItem(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("Content HappyPath", func(t *testing.T) {
		result, err := srv.UpdateListItemContent(context.Background(), 1, "updated-content")
		if err != nil {
			t.Fatalf("failed to update content: %s", err)
		}
		require.Equal(t, result.Content, "updated-content")
	})
	t.Run("Sort HappyPath", func(t *testing.T) {
		result, err := srv.UpdateListItemSort(context.Background(), 1, 2)
		if err != nil {
			t.Fatalf("failed to update sort: %s", err)
		}
		require.Equal(t, result.Sort, int64(2))
	})
	t.Run("Checked HappyPath,", func(t *testing.T) {
		result, err := srv.UpdateListItemChecked(context.Background(), 1, true)
		if err != nil {
			t.Fatalf("failed to update checked: %s", err)
		}
		require.Equal(t, result.Checked, true)
	})
}
func TestDeleteListItem(t *testing.T) {
	srv, closer, err := setupService()
	if err != nil {
		t.Fatalf("failed to setup service: %s", err)
	}
	defer closer.Close()
	t.Run("HappyPath", func(t *testing.T) {
		itemResult, err := srv.FetchListItem(context.Background(), 1)
		if err != nil {
			t.Fatalf("failed to fetch list item: %s", err)
		}
		deleteErr := srv.DeleteListItem(context.Background(), itemResult.ID)
		if deleteErr != nil {
			t.Fatalf("failed to delete item: %s", err)
		}
		listResult, err := srv.FetchAllItemsFromList(context.Background(), 1, 10)
		if err != nil {
			t.Fatalf("failed to fetch items from list: %s", err)
		}
		require.NotContains(t, listResult, itemResult)
	})
}
