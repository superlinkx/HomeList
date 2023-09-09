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
