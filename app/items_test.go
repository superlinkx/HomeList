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

package app_test

import (
	"testing"

	"github.com/superlinkx/HomeList/model"
)

var (
	successNoOffsetItems = []model.Item{
		{ID: 0, ListID: 0, Content: "Item 0", Checked: false, Sort: 1024},
		{ID: 1, ListID: 0, Content: "Item 1", Checked: false, Sort: 2048},
		{ID: 2, ListID: 0, Content: "Item 2", Checked: false, Sort: 3072},
		{ID: 3, ListID: 0, Content: "Item 3", Checked: false, Sort: 4096},
		{ID: 4, ListID: 0, Content: "Item 4", Checked: false, Sort: 5120},
		{ID: 5, ListID: 0, Content: "Item 5", Checked: false, Sort: 6144},
		{ID: 6, ListID: 0, Content: "Item 6", Checked: false, Sort: 7168},
		{ID: 7, ListID: 0, Content: "Item 7", Checked: false, Sort: 8192},
		{ID: 8, ListID: 0, Content: "Item 8", Checked: false, Sort: 9216},
		{ID: 9, ListID: 0, Content: "Item 9", Checked: false, Sort: 10240},
	}

	successOffset5Items = []model.Item{
		{ID: 5, ListID: 0, Content: "Item 5", Checked: false, Sort: 6144},
		{ID: 6, ListID: 0, Content: "Item 6", Checked: false, Sort: 7168},
		{ID: 7, ListID: 0, Content: "Item 7", Checked: false, Sort: 8192},
		{ID: 8, ListID: 0, Content: "Item 8", Checked: false, Sort: 9216},
		{ID: 9, ListID: 0, Content: "Item 9", Checked: false, Sort: 10240},
		{ID: 10, ListID: 0, Content: "Item 10", Checked: false, Sort: 11264},
		{ID: 11, ListID: 0, Content: "Item 11", Checked: false, Sort: 12288},
		{ID: 12, ListID: 0, Content: "Item 12", Checked: false, Sort: 13312},
		{ID: 13, ListID: 0, Content: "Item 13", Checked: false, Sort: 14336},
		{ID: 14, ListID: 0, Content: "Item 14", Checked: false, Sort: 15360},
	}
)

func TestApp_AllItems(t *testing.T) {}

func TestApp_GetItem(t *testing.T) {}

func TestApp_CreateItem(t *testing.T) {}

func TestApp_UpdateItem(t *testing.T) {}

func TestApp_DeleteItem(t *testing.T) {}