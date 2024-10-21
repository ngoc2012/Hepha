FROM alpine:latest

RUN apk update && \
    apk add --no-cache nginx \
    && rm -rf /var/cache/apk/*

RUN mkdir -p /var/www/html
RUN chown -R nginx:nginx /var/www/html
RUN chmod -R 755 /var/www/html

# Copy your custom Nginx configuration file to the container
COPY nginx.conf /etc/nginx/http.d
EXPOSE 3001

# Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]