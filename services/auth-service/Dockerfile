FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./


COPY services/common ./services/common
COPY services/auth-service ./services/auth-service

RUN go mod download

WORKDIR /app/services/auth-service
RUN go build -o auth-service main.go

# Final image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/services/auth-service/auth-service .
EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=2s --start-period=5s \
  CMD wget -qO- http://localhost:8080/health || exit 1
  
CMD ["./auth-service"]