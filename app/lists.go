package app

import (
	"context"
	"errors"

	"github.com/superlinkx/HomeList/data/adapter"
	"github.com/superlinkx/HomeList/model"
)

func (a App) AllLists(ctx context.Context, limit int32, offset int32) ([]model.List, error) {
	if ls, err := a.adapter.AllLists(ctx, limit, offset); errors.Is(err, adapter.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, ErrInternal
	} else {
		return ls, nil
	}
}

func (a App) GetList(ctx context.Context, id int64) (model.List, error) {
	return model.List{}, ErrInternal
}

func (a App) CreateList(ctx context.Context, name string) (model.List, error) {
	return model.List{}, ErrInternal
}

func (a App) UpdateList(ctx context.Context, id int64, name string) (model.List, error) {
	return model.List{}, ErrInternal
}

func (a App) DeleteList(ctx context.Context, id int64) error {
	return ErrInternal
}
