#!/bin/sh

if [ "$POSTGRES_INIT" = "true" ]; then
  postgres -c config_file=/var/lib/postgresql/data/postgresql.conf -D /var/lib/postgresql/data -f /app/create_db.sql
else
  postgres -c config_file=/var/lib/postgresql/data/postgresql.conf -D /var/lib/postgresql/data
fi

# psql -h localhost -U postgres -d mydatabase
