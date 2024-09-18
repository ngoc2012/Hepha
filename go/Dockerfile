FROM alpine:latest AS builder

RUN apk update && \
    apk add --no-cache make \
    && rm -rf /var/cache/apk/*
    
RUN mkdir -p /var/lib/files

WORKDIR /app

COPY go.sh .

ENV PORT $GO_PORT
RUN chmod +x go.sh

EXPOSE $GO_PORT

ENTRYPOINT [ "./go.sh" ]

# Stage 2: Create a smaller image with only the binary
FROM alpine:latest

RUN apk update && \
    apk add --no-cache make \
    && rm -rf /var/cache/apk/*
    
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /root/go/bin/air .

# Command to run the executable
CMD ["./air", "-c", "air.toml"]