FROM alpine:latest

RUN apk update && \
    apk add --no-cache mysql mysql-client expect \
    && rm -rf /var/cache/apk/*

RUN addgroup hefa
RUN adduser -D -G hefa hefa

WORKDIR /app

COPY mysql.sh .
COPY create_db.sql .
RUN chmod +x mysql.sh

EXPOSE 3306

ENTRYPOINT [ "./mysql.sh" ]