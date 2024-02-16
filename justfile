# MIT License

# Copyright (c) 2023 Alyx Holms
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

_default:
	@just --list --unsorted

export HL_HOST_URL := env_var_or_default("HL_HOST_URL", ":2552")

init:
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.24.0
  go install github.com/rubenv/sql-migrate/...@v1.5.2
  go install github.com/cosmtrek/air@v1.44.0
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
  go install github.com/superlinkx/go-runner@latest
  go install github.com/vektra/mockery/v2@v2.38.0
  go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
  just tidy

# Run application in development mode
dev *ARGS='':
  @air -c .air.toml -- {{ARGS}}

# Run application in debug mode
debug *ARGS='':
  @air -c .air.debug.toml -- {{ARGS}}

# Tidy module
tidy:
  @go mod tidy

# Codegen
codegen:
  @sqlc generate
  @go generate ./...
  @go-runner license
