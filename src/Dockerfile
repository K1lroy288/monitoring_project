FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o monitoring_service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/monitoring_service .

EXPOSE 3425

CMD ["./monitoring_service"]