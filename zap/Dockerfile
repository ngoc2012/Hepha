FROM alpine:latest

RUN apk update && \
  	apk add --no-cache python3 \
	lsof \
	xz \
	# supervisor \
	&& rm -rf /var/lib/apt/lists/*
# lsof check if files are open
# supervisor is used to run multiple processes (restart the web server)

# RUN mkdir -p /var/lib/files

RUN adduser -D hefa
RUN mkdir -p /home/hefa/app
RUN chown -R hefa:hefa /home/hefa/app

WORKDIR /app

COPY zap.sh .
COPY rebuild.sh .
# COPY zig-linux-x86_64-0.13.0.tar.xz .
# COPY supervisord.conf /etc/supervisor/conf.d/zap.conf

# RUN mkdir -p /app/zig
# RUN xz -d zig-linux-x86_64-0.13.0.tar.xz && tar -xvf zig-linux-x86_64-0.13.0.tar -C /app

RUN chmod +x zap.sh
RUN chmod +x rebuild.sh

# USER hefa

# ENV PATH=/app/zig-linux-x86_64-0.13.0:$PATH



# RUN zig build
#
#WORKDIR /home/hefa/app/build/zig-out/bin

EXPOSE 4000

#CMD [ "zap && tail -f /dev/null" ]
# CMD [ "tail", "-f", "/dev/null" ]
ENTRYPOINT [ "./zap.sh" ]
