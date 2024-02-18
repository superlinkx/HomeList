package app

import (
	"context"

	"github.com/superlinkx/HomeList/model"
)

func (a App) AllItemsFromList(ctx context.Context, listID int64, limit int32, offset int32) ([]model.Item, error) {
	return []model.Item{}, ErrInternal
}

func (a App) GetItemFromList(ctx context.Context, listID int64, itemID int64) (model.Item, error) {
	return model.Item{}, ErrInternal
}

func (a App) CreateItem(ctx context.Context, listID int64, content string) (model.Item, error) {
	return model.Item{}, ErrInternal
}

func (a App) UpdateItem(ctx context.Context, listID int64, itemID int64, content string, checked bool, sort int64) (model.Item, error) {
	return model.Item{}, ErrInternal
}

func (a App) DeleteItem(ctx context.Context, listID int64, itemID int64) error {
	return ErrInternal
}
