#!/bin/sh

# Here's the order of operations when a container is started:

# The container is created based on the image specified in the Dockerfile.
# The volumes defined in the docker-compose.yml file are mounted to the container's file system.
# The container's network is set up.
# The container's environment variables are set.
# The container's working directory is set.
# The ENTRYPOINT and CMD instructions from the Dockerfile are executed.

chown -R hepha:hepha /var/lib/postgresql/data

su - hepha

# Initialize the PostgreSQL database
/usr/bin/initdb -D /var/lib/postgresql/data

if [ "$POSTGRES_INIT" = "true" ]; then
  postgres -c config_file=/app/postgresql.conf -D /var/lib/postgresql/data -f /app/create_db.sql
else
  postgres -c config_file=/app/postgresql.conf -D /var/lib/postgresql/data
fi

# psql -h localhost -U postgres -d mydatabase
