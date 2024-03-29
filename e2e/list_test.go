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

package e2e_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superlinkx/HomeList/oapiclient"
)

func TestGetLists(t *testing.T) {
	var (
		limit  int32 = 10
		offset int32 = 0
		params       = oapiclient.GetListsParams{
			Limit:  &limit,
			Offset: &offset,
		}
	)

	client, err := oapiclient.NewClientWithResponses("http://localhost:2552/api/v1")
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.GetListsWithResponse(context.Background(), &params)
	require.NoError(t, err)
	require.NotNil(t, resp.JSON200)

	rawbody := *resp.JSON200
	require.NotNil(t, rawbody)

	assert.Len(t, rawbody, 10)

	limit = 1

	resp, err = client.GetListsWithResponse(context.Background(), &params)
	require.NoError(t, err)
	require.NotNil(t, resp.JSON200)

	rawbody = *resp.JSON200
	require.NotNil(t, rawbody)

	assert.Len(t, rawbody, 1)
}
