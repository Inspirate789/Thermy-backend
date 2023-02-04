#!/bin/bash

docker build -t web-server test01/server/
docker run --name test-backend -p 8080:8080 web-server
