FROM alpine:latest
    
RUN mkdir -p /var/lib/files

WORKDIR /app

COPY go.sh .

ENV PORT $GO_PORT
RUN chmod +x go.sh

EXPOSE $GO_PORT

ENTRYPOINT [ "./go.sh" ]