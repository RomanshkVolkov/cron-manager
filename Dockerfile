# Developer: Romanshk Volkov - https://github.com/RomanshkVolkov
# Customer: Dwit México - https://dwitmexico.com
# Project: cron-manager

FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go build -o /cron-manager ./cmd/

# drop build enviroment
FROM alpine:latest

WORKDIR /app

COPY --from=builder /cron-manager .

COPY --from=builder ./static ./static

EXPOSE 8080

CMD ["./cron-manager"]
