#!/bin/sh

export PATH="/app/zig:$PATH"
export PATH="/app/zig:/app/build/zig-out/bin:$PATH"

python3 --version

zig version

cd build

if [ "$ZAP_BUILD" = "true" ]; then
    echo "Building zap..."
    zig build
fi

# echo "Running as user: $(whoami)"
# pwd
# ls /home/hefa/app/build/
# echo "==================="
# ls /home/hefa/app/build/zig-out
# echo "==================="
# ls /home/hefa/app/build/zig-out/bin
# echo "==================="

# chmod +x /home/hefa/app/build/zig-out/bin/zap

# su - hefa

# export PATH="/home/hefa/app/build/zig-out/bin:$PATH"

echo "Running as user: $(whoami)"

zap
# tail -f /dev/null