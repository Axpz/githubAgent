FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/agentServer cmd/server/main.go
RUN go build -o bin/agent cmd/agent/main.go

FROM alpine:latest

RUN apk add --no-cache bash

WORKDIR /app

COPY --from=builder /app/bin/agentServer /app/bin/agentServer
COPY --from=builder /app/bin/agent /app/bin/agent

CMD ["/app/bin/agentServer"]

