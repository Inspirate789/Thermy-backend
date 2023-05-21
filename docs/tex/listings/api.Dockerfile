# syntax=docker/dockerfile:1

FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/backend/main.go ./
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./swagger ./swagger

RUN go build -o /backend

FROM scratch

COPY --from=build /backend /backend
COPY --from=build /app/swagger ./swagger
COPY backend.env /

EXPOSE ${BACKEND_PORT}

ENTRYPOINT ["/backend", "-env=backend.env"]
