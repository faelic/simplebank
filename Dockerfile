# Build stage
FROM golang:1.26.4-bookworm AS builder

WORKDIR /app

# Copy dependency files first for better build caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

COPY app.env .

EXPOSE 8080

CMD ["/app/main"]