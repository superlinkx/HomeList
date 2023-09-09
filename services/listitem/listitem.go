package listitem

import (
	"context"
	"fmt"

	"github.com/superlinkx/HomeList/db/sqlite"
)

type Service struct {
	queries *sqlite.Queries
}

type ListItem struct {
	ID      int64
	ListID  int64
	Content string
	Checked bool
	Sort    int64
}

func NewService(queries *sqlite.Queries) Service {
	return Service{queries: queries}
}

func (s Service) FetchListItem(ctx context.Context, id int64) (ListItem, error) {
	if listitem, err := s.queries.GetListItem(ctx, id); err != nil {
		return ListItem{}, fmt.Errorf("failed to fetch list item: %w", err)
	} else {
		return ListItem(listitem), nil
	}
}

func (s Service) FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]ListItem, error) {
	var params = sqlite.AllItemsFromListParams{ListID: listID, Limit: limit}

	if results, err := s.queries.AllItemsFromList(ctx, params); err != nil {
		return []ListItem{}, fmt.Errorf("failed to fetch list items: %w", err)
	} else {
		var listitems = make([]ListItem, 0, len(results))
		for _, result := range results {
			listitem := ListItem(result)
			listitems = append(listitems, listitem)
		}
		return listitems, nil
	}
}

func (s Service) AddItemToList(ctx context.Context, listID int64, content string, sort int64) (ListItem, error) {
	params := sqlite.CreateListItemParams{
		ListID:  listID,
		Content: content,
		Sort:    sort,
	}

	if listItem, err := s.queries.CreateListItem(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to create list item: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s Service) UpdateListItemContent(ctx context.Context, itemID int64, content string) (ListItem, error) {
	params := sqlite.UpdateListItemTextParams{
		ID:      itemID,
		Content: content,
	}

	if listItem, err := s.queries.UpdateListItemText(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item text: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s Service) UpdateListItemSort(ctx context.Context, itemID int64, sort int64) (ListItem, error) {
	params := sqlite.UpdateListItemSortParams{
		ID:   itemID,
		Sort: sort,
	}

	if listItem, err := s.queries.UpdateListItemSort(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item sort: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s Service) UpdateListItemChecked(ctx context.Context, itemID int64, checked bool) (ListItem, error) {
	params := sqlite.UpdateListItemCheckedParams{
		ID:      itemID,
		Checked: checked,
	}

	if listItem, err := s.queries.UpdateListItemChecked(ctx, params); err != nil {
		return ListItem{}, fmt.Errorf("failed to update list item checked state: %w", err)
	} else {
		return ListItem(listItem), nil
	}
}

func (s Service) DeleteListItem(ctx context.Context, itemID int64) error {
	if err := s.queries.DeleteListItem(ctx, itemID); err != nil {
		return fmt.Errorf("failed to delete list item: %w", err)
	} else {
		return nil
	}
}
