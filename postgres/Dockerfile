# Use the official Alpine image as a base image
FROM alpine:latest

# Install PostgreSQL
RUN apk update && \
    apk add --no-cache postgresql

# Create a directory for PostgreSQL data

# RUN mkdir -p /var/lib/postgresql/data && \
#     chown -R hepha:hepha /var/lib/postgresql/data

# Set environment variables for PostgreSQL
# ENV POSTGRES_US/ER=postgres
# ENV POSTGRES_PASSWORD=mysecretpassword
# ENV POSTGRES_DB=mydatabase

RUN addgroup hepha
RUN adduser -D -G hepha hepha

WORKDIR /app

COPY postgres.sh .
COPY create_db.sql .
RUN chmod +x postgres.sh
COPY postgresql.conf .

# COPY init.sql /docker-entrypoint-initdb.d/init.sql

# Expose the PostgreSQL port
EXPOSE 5432

# RUN mkdir -p /var/lib/postgresql/data
# RUN chown postgres /var/lib/postgresql/data
RUN chown -R hepha:hepha /var/lib/postgresql/data

USER hepha

# Initialize the PostgreSQL database
RUN /usr/bin/initdb -D /var/lib/postgresql/data

# Start PostgreSQL server with the custom configuration file
# CMD ["postgres", "-c", "config_file=/var/lib/postgresql/data/postgresql.conf", "-D", "/var/lib/postgresql/data"]

ENTRYPOINT [ "./postgres.sh" ]