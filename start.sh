#!/bin/sh

set -e

# Create master.key file from environment variable
if [ ! -z "$MASTER_KEY" ]; then
  echo "Creating master.key file from environment variable"
  echo "$MASTER_KEY" | base64 -d > /app/config/master.key
  chmod 600 /app/config/master.key
fi

echo "run db migration"
goose --dir=db/migrate -allow-missing postgres $DB_URI up

echo "start server"
exec "$@"