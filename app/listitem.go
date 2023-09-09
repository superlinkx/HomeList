package app

import (
	"context"
	"fmt"
)

type ListItem struct {
	ID      int64
	ListID  int64
	Content string
	Checked bool
	Sort    int64
}

func (a Application) FetchListItem(ctx context.Context, itemID int64) (ListItem, error) {
	if result, err := a.services.ListItem.FetchListItem(ctx, itemID); err != nil {
		return ListItem{}, fmt.Errorf("application failed to fetch list item: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (a Application) FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]ListItem, error) {
	if results, err := a.services.ListItem.FetchAllItemsFromList(ctx, listID, limit); err != nil {
		return []ListItem{}, fmt.Errorf("application failed to fetch all items from list with id %d: %w", listID, err)
	} else {
		var listitems = make([]ListItem, 0, len(results))
		for _, result := range results {
			listitem := ListItem{
				ID:      result.ID,
				ListID:  result.ListID,
				Content: result.Content,
				Checked: result.Checked,
				Sort:    result.Sort,
			}
			listitems = append(listitems, listitem)
		}
		return listitems, nil
	}
}

func (a Application) AddItemToList(ctx context.Context, listID int64, content string, sort int64) (ListItem, error) {
	if result, err := a.services.ListItem.AddItemToList(ctx, listID, content, sort); err != nil {
		return ListItem{}, fmt.Errorf("application failed to add item to list: %w", err)
	} else {
		return ListItem{
			ID:      result.ID,
			ListID:  result.ListID,
			Content: result.Content,
			Checked: result.Checked,
			Sort:    result.Sort,
		}, nil
	}
}

func (a Application) UpdateListItemContent(ctx context.Context, itemID int64, content string) (ListItem, error) {
	if result, err := a.services.ListItem.UpdateListItemContent(ctx, itemID, content); err != nil {
		return ListItem{}, fmt.Errorf("application failed to update list item text: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (a Application) UpdateListItemChecked(ctx context.Context, itemID int64, checked bool) (ListItem, error) {
	if result, err := a.services.ListItem.UpdateListItemChecked(ctx, itemID, checked); err != nil {
		return ListItem{}, fmt.Errorf("application failed to update list item checked: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (a Application) UpdateListItemSort(ctx context.Context, itemID int64, sort int64) (ListItem, error) {
	if result, err := a.services.ListItem.UpdateListItemSort(ctx, itemID, sort); err != nil {
		return ListItem{}, fmt.Errorf("application failed to update list item sort: %w", err)
	} else {
		return ListItem(result), nil
	}
}

func (a Application) DeleteListItem(ctx context.Context, itemID int64) error {
	if err := a.services.ListItem.DeleteListItem(ctx, itemID); err != nil {
		return fmt.Errorf("application failed to delete list item: %w", err)
	} else {
		return nil
	}
}
