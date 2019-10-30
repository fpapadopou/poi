#!/usr/bin/env bash

set -e

echo "" > coverage.txt

# Run migrations
go run -tags psql ./migrations/migrate.go init
go run -tags psql ./migrations/migrate.go up

# Run tests
go test -tags psql -coverprofile=profile.out -coverpkg=./... ./...

cat profile.out
cat profile.out >> coverage.txt
rm profile.out
