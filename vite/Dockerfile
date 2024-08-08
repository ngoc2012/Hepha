FROM alpine:latest

RUN apk update && \
  	apk add --no-cache nodejs npm && \
	rm -rf /var/lib/apt/lists/*

WORKDIR /app

# COPY entrypoint.sh .
# RUN chmod +x entrypoint.sh
# ENTRYPOINT [ "./entrypoint.sh" ]

EXPOSE $VITE_PORT

CMD [ "npm", "run", "dev" ]
# CMD ["npm", "start"]