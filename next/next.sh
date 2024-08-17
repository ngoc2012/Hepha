#!/bin/sh
# export ZAP_URL=$(busybox ping -c 1 zap | busybox head -1 | busybox awk -F'[()]' '{print $2}')

npm install --legacy-peer-deps
npm run build

exec "$@"