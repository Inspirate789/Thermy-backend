#!/bin/bash

docker build -t web-client test01/js-client/
docker run --name test-frontend -p 3000:3000 -i -t web-client
