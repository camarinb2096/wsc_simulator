FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o wsc-simulator cmd/main.go

EXPOSE 80

CMD ["./wsc-simulator"]

