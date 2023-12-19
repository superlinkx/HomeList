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

package restapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/HomeList/handler"
	"github.com/superlinkx/HomeList/restapi"
)

func TestNewServer(t *testing.T) {
	var (
		fullConfig = restapi.Config{
			HostURL: "localhost:8080",
		}
		emptyConfig = restapi.Config{}
	)

	t.Run("full config", func(t *testing.T) {
		s := restapi.NewServer(fullConfig, handler.Handlers{})
		assert.Equal(t, "localhost:8080", s.Addr)
	})

	t.Run("empty config", func(t *testing.T) {
		s := restapi.NewServer(emptyConfig, handler.Handlers{})
		assert.Equal(t, ":80", s.Addr)
	})
}
