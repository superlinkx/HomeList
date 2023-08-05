package app

import (
	"context"
	"fmt"

	"github.com/superlinkx/HomeList/services"
)

type Application struct {
	services services.Services
}

type List struct {
	ID   int64
	Name string
}

func NewApplication(services services.Services) Application {
	return Application{
		services: services,
	}
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
