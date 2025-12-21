# Use official Go image as build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o auth-service main.go

# Final lightweight image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/auth-service .
COPY .env.example .env
EXPOSE 8080

# Health endpoint for K8s
HEALTHCHECK --interval=10s --timeout=2s --start-period=5s \
  CMD wget -qO- http://localhost:8080/health || exit 1
  
CMD ["./auth-service"]