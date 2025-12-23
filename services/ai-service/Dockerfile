FROM golang:1.25-alpine AS builder
WORKDIR /app


COPY go.mod go.sum ./
COPY services/common ./services/common
COPY services/ai-service ./services/ai-service

RUN go mod download

WORKDIR /app/services/ai-service
RUN go build -o ai-service cmd/main.go

# Final image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/services/ai-service/ai-service .
EXPOSE 8082

CMD ["./ai-service"]