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
	} else {
		defer closer.Close()

		if li, err := srv.FetchListItem(context.Background(), 1); err != nil {
			t.Fatalf("failed to fetch list item: %s", err)
		} else {
			require.Equal(t, int64(1), li.ID)
			require.Equal(t, int64(1), li.ListID)
			require.Equal(t, int64(1024), li.Sort)
			require.Equal(t, "Hello world!", li.Content)
			require.Equal(t, false, li.Checked)
		}
	}
}
