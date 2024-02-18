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
)

func (s App) AllItemsFromList(ctx context.Context, listID int64, limit int32, offset int32) ([]model.Item, error) {
	return []model.Item{}, errors.New("not implemented")
}

func (s App) GetItemFromList(ctx context.Context, listID int64, itemID int64) (model.Item, error) {
	return model.Item{}, errors.New("not implemented")
}

func (s App) CreateItem(ctx context.Context, listID int64, content string) (model.Item, error) {
	return model.Item{}, errors.New("not implemented")
}

func (s App) UpdateItem(ctx context.Context, listID int64, itemID int64, content string, checked bool, sort int64) (model.Item, error) {
	return model.Item{}, errors.New("not implemented")
}

func (s App) DeleteItem(ctx context.Context, listID int64, itemID int64) error {
	return errors.New("not implemented")
}
