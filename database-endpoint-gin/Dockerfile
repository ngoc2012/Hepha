FROM alpine:latest
    
RUN mkdir -p /var/lib/files

WORKDIR /app

COPY gin.sh .

ENV PORT $GIN_PORT
RUN chmod +x gin.sh

EXPOSE $GIN_PORT

ENTRYPOINT [ "./gin.sh" ]