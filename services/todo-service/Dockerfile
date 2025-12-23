FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
COPY services/common ./services/common
COPY services/todo-service ./services/todo-service

RUN go mod download

WORKDIR /app/services/todo-service
RUN go build -o todo-service main.go

# Final image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/services/todo-service/todo-service .
EXPOSE 8081

HEALTHCHECK --interval=10s --timeout=2s --start-period=5s \
  CMD wget -qO- http://localhost:8081/health || exit 1
  
CMD ["./todo-service"]