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

package main

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/superlinkx/HomeList/bootstrap"
	"github.com/superlinkx/HomeList/environment"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -config oapiclient/.oapi-codegen.yaml -o oapiclient/oapiclient.gen.go ./docs/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -config oapiserver/.oapi-codegen.yaml -o oapiserver/oapiserver.gen.go ./docs/openapi.yaml
//go:generate go run github.com/vektra/mockery/v2@v2.41.0

func main() {
	if config, err := environment.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %s\n", err)
	} else if db, err := sql.Open("sqlite", ":memory:?_pragma=foreign_keys(1)"); err != nil {
		log.Fatalf("Failed to open connection to database: %s\n", err)
	} else if err := bootstrap.MigrateDB(db); err != nil {
		db.Close()
		log.Fatalf("Failed to migrate database: %s\n", err)
	} else {
		defer db.Close()
		bootstrap.StartApp(db, config)
	}
}
