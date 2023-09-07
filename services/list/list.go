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

func (s Service) GetList(ctx context.Context, id int64) (List, error) {
	if result, err := s.queries.GetList(ctx, id); err != nil {
		return List{}, fmt.Errorf("unable to get list: %w", err)
	} else {
		return List{ID: result.ID, Name: result.Name}, nil
	}
}

func (s Service) CreateList(ctx context.Context, name string) (List, error) {
	if result, err := s.queries.CreateList(ctx, name); err != nil {
		return List{}, fmt.Errorf("unable to create list: %w", err)
	} else {
		return List{ID: result.ID, Name: result.Name}, nil
	}
}

func (s Service) UpdateList(ctx context.Context, id int64, name string) (List, error) {
	if result, err := s.queries.RenameList(ctx, sqlite.RenameListParams{ID: id, Name: name}); err != nil {
		return List{}, fmt.Errorf("unable to update list: %w", err)
	} else {
		return List{ID: result.ID, Name: result.Name}, nil
	}
}

func (s Service) DeleteList(ctx context.Context, id int64) error {
	if err := s.queries.DeleteList(ctx, id); err != nil {
		return fmt.Errorf("unable to delete list: %w", err)
	} else {
		return nil
	}
}
