#!/bin/bash

docker pull node:10.23.1-alpine3.11
docker pull golang:1.17.12-alpine3.15
docker-compose -f test01/docker-compose-linux.yml up -d
