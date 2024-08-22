FROM alpine:latest
    
RUN mkdir -p /var/lib/files

WORKDIR /app

COPY gin.sh .
# COPY supervisord.conf /etc/supervisor/conf.d/zap.conf

ENV PORT $GIN_PORT
RUN chmod +x gin.sh
# RUN chmod +x rebuild.sh

EXPOSE $GIN_PORT

ENTRYPOINT [ "./gin.sh" ]