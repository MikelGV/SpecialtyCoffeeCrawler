FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./internal/server/main.go

FROM debian:stable-slim
WORKDIR /app

COPY --from=builder /app/server /app/server
COPY .env /app/.env

RUN chmod +x /app/server

EXPOSE 8080

CMD ["/app/server"]
