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
	migrate "github.com/rubenv/sql-migrate"
	"github.com/superlinkx/HomeList/api"
	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/environment"
	"github.com/superlinkx/HomeList/server"
	"github.com/superlinkx/HomeList/services"
)

func main() {
	if config, err := environment.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	} else if db, err := sql.Open("sqlite", ":memory:?_pragma=foreign_keys(1)"); err != nil {
		log.Fatalf("Failed to open connection to database: %s", err)
	} else {
		defer db.Close()

		migrations := &migrate.FileMigrationSource{
			Dir: "sql/migrations",
		}

		if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
			log.Fatalf("Failed to migrate database: %s", err)
		}

		s := services.Init(db)
		application := app.NewApplication(s)
		api := api.NewAPI(application)
		server.StartServer(server.Config{HostURL: config.HostURL}, api)
	}
}
