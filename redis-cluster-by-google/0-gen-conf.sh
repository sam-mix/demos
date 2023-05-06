#!/usr/bin/env bash

set -e

mkdir -p conf && \
for port in `seq 6371 6376`; do \
  PORT=${port} HOST=192.168.0.219 envsubst < redis-cluster.tmpl > conf/redis_${port}.conf; \
done
