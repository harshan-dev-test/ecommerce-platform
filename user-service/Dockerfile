FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8080

ENTRYPOINT [ "./user-service" ]