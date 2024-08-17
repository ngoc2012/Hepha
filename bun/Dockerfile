FROM alpine:latest

RUN apk update && \
  	apk add --no-cache curl \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY bun.sh .
RUN chmod +x bun.sh
ENTRYPOINT [ "./bun.sh" ]

EXPOSE 3000

CMD [ "npm", "run", "dev" ]
# CMD ["npm", "start"]