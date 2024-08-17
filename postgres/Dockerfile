# Use the official Alpine image as a base image
FROM alpine:latest

# Install PostgreSQL
RUN apk update && \
    apk add --no-cache postgresql \
    bash su-exec \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY postgres.sh .
COPY create_db.sql .
RUN chmod +x postgres.sh
COPY postgresql.conf .

RUN mkdir /run/postgresql
RUN chown postgres:postgres /run/postgresql/

# RUN addgroup postgres
# RUN adduser -D -G postgres hepha

# RUN mkdir /run/postgresql && \
#     chown hepha:postgres /run/postgresql && \
#     chmod 755 /run/postgresql

# Expose the PostgreSQL port
EXPOSE $POSTGRES_PORT

# Start PostgreSQL server with the custom configuration file
# CMD ["postgres", "-c", "config_file=/var/lib/postgresql/data/postgresql.conf", "-D", "/var/lib/postgresql/data"]

ENTRYPOINT [ "./postgres.sh" ]
# CMD [ "tail", "-f", "/dev/null" ]