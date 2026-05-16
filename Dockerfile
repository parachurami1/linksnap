# Stage 1 — Build
FROM golang:1.26.3-alpine3.23 AS builder

WORKDIR /app

# Copy dependency files first (better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o linksnap ./main.go


# Stage 2 — Final image
FROM alpine:latest

WORKDIR /app

# Copy only the binary from the builder
COPY --from=builder /app/linksnap .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./linksnap"]