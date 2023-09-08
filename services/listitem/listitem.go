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
		return ListItem{
			ID:      listitem.ID,
			ListID:  listitem.ListID,
			Content: listitem.Content,
			Checked: listitem.Checked,
			Sort:    listitem.Sort,
		}, nil
	}
}

func (s Service) FetchAllItemsFromList(ctx context.Context, listID int64, limit int64) ([]ListItem, error) {
	var params = sqlite.AllItemsFromListParams{ListID: listID, Limit: limit}

	if results, err := s.queries.AllItemsFromList(ctx, params); err != nil {
		return nil, fmt.Errorf("failed to fetch list items: %w", err)
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
