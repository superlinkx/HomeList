#!/usr/bin/env bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.20.0
go install github.com/rubenv/sql-migrate/...@v1.5.2
go install github.com/cosmtrek/air@v1.44.0
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
go mod init