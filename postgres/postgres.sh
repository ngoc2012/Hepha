#!/bin/sh

# https://public.dalibo.com/exports/formation/manuels/modules/b/b.handout.html

# Here's the order of operations when a container is started:

# The container is created based on the image specified in the Dockerfile.
# The volumes defined in the docker-compose.yml file are mounted to the container's file system.
# The container's network is set up.
# The container's environment variables are set.
# The container's working directory is set.
# The ENTRYPOINT and CMD instructions from the Dockerfile are executed.

# echo "Running as user: $(whoami)"

# id

chown postgres:postgres /var/lib/postgresql/data
chmod 0700 /var/lib/postgresql/data
# RUN initdb -D /var/lib/postgresql/data

# The reason why 0700 is used instead of 700 is to ensure that there are no special permissions set for the owner. The leading 0 in 0700 explicitly sets the special permissions to none, while 700 would leave any existing special permissions unchanged.
# Special permissions include the setuid, setgid, and sticky bits. The setuid bit allows a user to run a program with the privileges of the owner of the program, while the setgid bit allows a user to run a program with the privileges of the group owner of the program. The sticky bit prevents users from deleting or renaming files in a directory, even if they have write permission for the directory.

# Switch to the postgres user
exec su-exec postgres bash -c "
echo 'Running as user: $(whoami)'

# Initialize the PostgreSQL database if not already initialized
if [ -z \"\$(ls -A /var/lib/postgresql/data)\" ]; then
  echo 'Initializing database...'
  /usr/bin/initdb -D /var/lib/postgresql/data
  echo \"host all all 0.0.0.0/0 md5\" >> /var/lib/postgresql/data/pg_hba.conf
  echo \"listen_addresses='*'\" >> /var/lib/postgresql/data/postgresql.conf

  pg_ctl start -D /var/lib/postgresql/data -l /var/lib/postgresql/data/logfile

  echo 'Create tables.'
  psql -U postgres -d postgres -a -f /app/create_db.sql

  echo \"Creating user $POSTGRES_USER...\"
  psql -v ON_ERROR_STOP=1 --username postgres --dbname \"$POSTGRES_DB\" --host \"localhost\" --port \"$POSTGRES_PORT\" <<-EOSQL
      CREATE USER $POSTGRES_USER WITH ENCRYPTED PASSWORD '$POSTGRES_PASSWORD';
      GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
      GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO $POSTGRES_USER;

EOSQL

  
  pg_ctl stop -D /var/lib/postgresql/data
else
  echo 'Database already initialized.'
fi

# Start PostgreSQL
postgres -D /var/lib/postgresql/data
"

