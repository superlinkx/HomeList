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

package integration

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations
var migrationFS embed.FS

func ConnectDatabase() (*sql.DB, error) {
	if db, err := sql.Open("sqlite", ":memory:?_pragma=foreign_keys(1)"); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func ResetDatabase(db *sql.DB) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationFS,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Down); err != nil {
		return fmt.Errorf("failed to migrate down: %w", err)
	} else if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return fmt.Errorf("failed to migrate up: %w", err)
	} else {
		return nil
	}
}
