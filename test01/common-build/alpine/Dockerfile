FROM golang:1.17.12-alpine3.15

WORKDIR /app/
ENV GO111MODULE=auto
RUN apk --no-cache add --virtual build-deps
RUN apk update && apk add --no-cache protobuf-dev~=3.18.1-r1 protoc~=3.18.1-r1 bash~=5.1.16-r0 git~=2.34.6-r0 && rm /var/cache/apk/*

# Install grpc
RUN go get -u -v github.com/golang/protobuf/proto
RUN go get -u -v github.com/golang/protobuf/protoc-gen-go
RUN go get -u -v google.golang.org/grpc

ENV PATH=$PATH:$GOPATH/bin:/opt/protoc/bin
