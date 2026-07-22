# Build stage
FROM golang:1.26.4-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates curl tar && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY db/migration ./db/migration
COPY start.sh .

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz | tar -xz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /app/start.sh

EXPOSE 8080

CMD ["/app/start.sh"]