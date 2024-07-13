#!/bin/sh

set -e
# echo 'run db migration'
# /usr/bin/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "start the app"
exec "$@"