#!/usr/bin/env bash

set -e

echo "" > coverage.txt

# Run migrations
cd migrations
go run migrate.go up
cd -

# Run tests
go test -tags psql -coverprofile=profile.out -coverpkg=./... ./...

cat profile.out
cat profile.out >> coverage.txt
rm profile.out
