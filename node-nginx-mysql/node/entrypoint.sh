#!/bin/sh

npm i

dockerize -wait tcp://db:3306 -timeout 30s

node /home/app/src/index.js
