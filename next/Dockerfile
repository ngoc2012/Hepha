FROM alpine:latest

RUN apk update && \
  	apk add --no-cache nodejs npm \
	curl \
	busybox \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY next.sh .
RUN chmod +x next.sh
ENTRYPOINT [ "./next.sh" ]

EXPOSE 3000

CMD [ "npm", "run", "dev" ]
# CMD ["npm", "start"]