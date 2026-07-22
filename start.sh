#!/bin/sh

set -e

echo "running database migrations..."
migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "starting app..."
exec /app/main