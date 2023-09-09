package app

import (
	"context"
	"fmt"
)

type List struct {
	ID   int64
	Name string
}

func (a Application) AllLists(ctx context.Context, limit int64) ([]List, error) {
	var lists = make([]List, 0, limit)
	if results, err := a.services.List.AllLists(ctx, limit); err != nil {
		return lists, fmt.Errorf("application failed enumerate lists: %w", err)
	} else {
		for _, result := range results {
			lists = append(lists, List{
				ID:   result.ID,
				Name: result.Name,
			})
		}

		return lists, nil
	}
}

func (a Application) GetList(ctx context.Context, id int64) (List, error) {
	if result, err := a.services.List.GetList(ctx, id); err != nil {
		return List{}, fmt.Errorf("application failed to get list with id %d: %w", id, err)
	} else {
		return List{ID: result.ID, Name: result.Name}, nil
	}
}

func (a Application) CreateList(ctx context.Context, name string) (List, error) {
	if result, err := a.services.List.CreateList(ctx, name); err != nil {
		return List{}, fmt.Errorf("application failed to create list: %w", err)
	} else {
		return List{ID: result.ID, Name: result.Name}, nil
	}
}

func (a Application) UpdateList(ctx context.Context, id int64, name string) (List, error) {
	if result, err := a.services.List.UpdateList(ctx, id, name); err != nil {
		return List{}, fmt.Errorf("application failed to update list: %w", err)
	} else {
		return List{ID: result.ID, Name: result.Name}, nil
	}
}

func (a Application) DeleteList(ctx context.Context, id int64) error {
	if err := a.services.List.DeleteList(ctx, id); err != nil {
		return fmt.Errorf("application failed to delete list: %w", err)
	} else {
		return nil
	}
}
