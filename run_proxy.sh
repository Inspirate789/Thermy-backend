#!/bin/bash

docker build -t web-proxy test01/envoy/
docker run --name test-proxy -p 8000:8000 web-proxy
