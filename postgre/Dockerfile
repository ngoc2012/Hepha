# Use the official Alpine image as a base image
FROM alpine:latest

# Install PostgreSQL
RUN apk update && \
    apk add --no-cache postgresql

# Create a directory for PostgreSQL data
RUN mkdir -p /var/lib/postgresql/data

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=mysecretpassword
ENV POSTGRES_DB=mydatabase

# Copy the custom postgresql.conf file into the container
COPY postgresql.conf /var/lib/postgresql/data/postgresql.conf

# Expose the PostgreSQL port
EXPOSE 5432

# Initialize the PostgreSQL database
RUN /usr/bin/initdb -D /var/lib/postgresql/data

# Start PostgreSQL server with the custom configuration file
# CMD ["postgres", "-c", "config_file=/var/lib/postgresql/data/postgresql.conf", "-D", "/var/lib/postgresql/data"]

ENTRYPOINT [ "./postgres.sh" ]