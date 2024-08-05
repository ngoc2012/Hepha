FROM alpine:latest
    
RUN mkdir -p /var/lib/files

WORKDIR /app

COPY gin.sh .
# COPY supervisord.conf /etc/supervisor/conf.d/zap.conf

RUN chmod +x gin.sh
# RUN chmod +x rebuild.sh

EXPOSE 8080

ENTRYPOINT [ "./gin.sh" ]