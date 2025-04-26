#!/usr/bin/env bash
set -e
if [ -z "$MONGO_URI" ]; then
  echo "MONGO_URI not set"
  exit 1
fi
go run cmd/migrate/*.go --uri "$MONGO_URI" --path "./migrations"