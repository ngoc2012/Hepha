FROM alpine:latest

RUN apk update && \
  	apk add --no-cache nodejs npm && \
	rm -rf /var/lib/apt/lists/*

WORKDIR /app

EXPOSE $REMIX_PORT

CMD [ "npm", "run", "dev" ]