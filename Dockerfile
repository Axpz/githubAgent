FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/agentServer cmd/server/main.go
RUN go build -o bin/agent cmd/agent/main.go

FROM alpine:latest

RUN apk add --no-cache bash curl

RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        curl -LO "https://dl.k8s.io/release/v1.32.0/bin/linux/amd64/kubectl"; \
    elif [ "$ARCH" = "aarch64" ]; then \
        curl -LO "https://dl.k8s.io/release/v1.32.0/bin/linux/arm64/kubectl"; \
    else \
        echo "Unsupported architecture"; exit 1; \
    fi && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin/

WORKDIR /app

COPY --from=builder /app/bin/agentServer /app/bin/agentServer
COPY --from=builder /app/bin/agent /app/bin/agent

CMD ["/app/bin/agentServer"]

