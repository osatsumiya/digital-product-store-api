FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP_DIR=.
RUN go build -o app ${APP_DIR}

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080
EXPOSE 8082
EXPOSE 8083

CMD ["./app"]