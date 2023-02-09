# # syntax=docker/dockerfile:1

# ## Build
# FROM golang:1.20.0-alpine3.17 AS build

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

# RUN go build -o gin-test

# ## Deploy
# FROM alpine:3.17

# WORKDIR /

# COPY --from=build gin-test gin-test

# EXPOSE 8080

# # USER nonroot:nonroot

# CMD ["./gin-test"]

# syntax=docker/dockerfile:1

FROM golang:1.20.0-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]