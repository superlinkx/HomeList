// MIT License
//
// Copyright (c) 2023 Alyx Holms
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package app

import (
	"context"
	"errors"

	"github.com/superlinkx/HomeList/app/model"
	"github.com/superlinkx/HomeList/data/adapter"
)

func (s App) AllLists(ctx context.Context, limit int32, offset int32) ([]model.List, error) {
	if ls, err := s.adapter.AllLists(ctx, limit, offset); errors.Is(err, adapter.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, ErrInternal
	} else {
		return ls, nil
	}
}

func (s App) GetList(ctx context.Context, id int64) (model.List, error) {
	return model.List{}, errors.New("not implemented")
}

func (s App) CreateList(ctx context.Context, name string) (model.List, error) {
	return model.List{}, errors.New("not implemented")
}

func (s App) UpdateList(ctx context.Context, id int64, name string) (model.List, error) {
	return model.List{}, errors.New("not implemented")
}

func (s App) DeleteList(ctx context.Context, id int64) error {
	return errors.New("not implemented")
}
