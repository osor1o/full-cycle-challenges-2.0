#!/bin/sh
dockerize -wait tcp://app:3000 -timeout 30s

exec "$@"
