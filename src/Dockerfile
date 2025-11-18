FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN GOOS=linux go build -o /monitoring_service

EXPOSE 3425

CMD ["/monitoring_service"]