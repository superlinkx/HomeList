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

package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/data/sqlite/adapter"
	"github.com/superlinkx/HomeList/environment"
	"github.com/superlinkx/HomeList/handler"
	"github.com/superlinkx/HomeList/restapi"
)

func MigrateDB(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "sql/sqlite/migrations",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return fmt.Errorf("migrate up failed: %w", err)
	}

	return nil
}

func StartApp(db *sql.DB, config environment.Config) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	dataAdapter := adapter.NewSqliteAdapter(db)
	app := app.NewApp(dataAdapter)
	handlers := handler.NewHandlers(app)
	srv, err := restapi.NewServer(restapi.Config{HostURL: config.HostURL}, handlers)
	if err != nil {
		log.Fatalf("Failed to create server: %s\n", err)
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failure: %s\n", err)
		}
	}()

	log.Printf("Server is running on %s", srv.Addr)

	<-done
	log.Println("Application stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %s\n", err)
	}

	log.Println("Server exited cleanly")
}
