#!/bin/sh

export PATH="/home/hepha/app/zig:/home/hepha/app/build/zig-out/bin:$PATH"

zig version

cd build

if [ "$ZAP_BUILD" = "true" ]; then
    echo "Building zap..."
    zig build
fi

# /usr/bin/supervisord -c /etc/supervisor/conf.d/zap.conf

# supervisorctl restart zap