#!/bin/sh

set -e

# Create master.key file from environment variable
if [ ! -z "$MASTER_KEY" ]; then
  echo "Creating master.key file from environment variable"
  echo "$MASTER_KEY" | base64 -d > /app/config/master.key
  chmod 600 /app/config/master.key
fi

echo "start server"
exec "$@"