#!/usr/bin/env bash
set -euo pipefail

echo "Waiting for Postgres at ${DATABASE_URL:-unset}..."
sleep 2

if [ -d "/app/backend/migrations" ]; then
  echo "Running goose up..."
  goose -dir /app/backend/migrations postgres "$DATABASE_URL" up || true
fi

go mod tidy

# Стартуем live-reload
echo "Starting air..."
cd /app/backend
exec air -c /app/backend/.air.toml