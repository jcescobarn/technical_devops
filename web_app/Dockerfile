FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /usr/local/bin

RUN adduser -D appuser

COPY --from=builder /app/app .

RUN chmod +x app

USER appuser

ENTRYPOINT ["./app"]