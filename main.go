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

//go:generate mockery

package main

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

	_ "github.com/glebarez/go-sqlite"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/superlinkx/HomeList/app"
	"github.com/superlinkx/HomeList/environment"
	"github.com/superlinkx/HomeList/handlers"
	"github.com/superlinkx/HomeList/restapi"
	"github.com/superlinkx/HomeList/services"
)

func main() {
	if config, err := environment.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %s\n", err)
	} else if db, err := sql.Open("sqlite", ":memory:?_pragma=foreign_keys(1)"); err != nil {
		log.Fatalf("Failed to open connection to database: %s\n", err)
	} else if err := migrateDB(db); err != nil {
		db.Close()
		log.Fatalf("Failed to migrate database: %s\n", err)
	} else {
		defer db.Close()
		startApp(db, config)
	}
}

func migrateDB(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "sql/migrations",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		return fmt.Errorf("migrate up failed: %w", err)
	}

	return nil
}

func startApp(db *sql.DB, config environment.Config) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	services := services.Init(db)
	application := app.NewApplication(services)
	handlers := handlers.NewHandlers(application)
	srv := restapi.NewServer(restapi.Config{HostURL: config.HostURL}, handlers)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failure: %s\n", err)
		}
	}()

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
