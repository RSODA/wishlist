FROM golang:1.25.7-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bin/sso cmd/main/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/sso .
COPY --from=builder /app/.env .env

CMD ["./sso"]