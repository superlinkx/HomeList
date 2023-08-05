package list

import (
	"context"
	"fmt"

	"github.com/superlinkx/HomeList/db/sqlite"
)

type Service struct {
	queries *sqlite.Queries
}

type List struct {
	ID   int64
	Name string
}

func NewService(queries *sqlite.Queries) Service {
	return Service{queries: queries}
}

func (s Service) AllLists(ctx context.Context, limit int64) ([]List, error) {
	var (
		lists = make([]List, 0, limit)
	)

	if results, err := s.queries.AllLists(ctx, limit); err != nil {
		return lists, fmt.Errorf("unable to get all lists: %w", err)
	} else if results == nil {
		return lists, nil
	} else {
		for _, result := range results {
			lists = append(lists, List{ID: result.ID, Name: result.Name})
		}
		return lists, nil
	}
}
