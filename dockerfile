# FROM golang:alpine AS builder

# RUN apk add --no-cache git

# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o wsc-simulator cmd/main.go

# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /app/wsc-simulator /app/wsc-simulator

# EXPOSE 80

# CMD ["./wsc-simulator"]

FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o wsc-simulator cmd/main.go

EXPOSE 80

CMD ["./wsc-simulator"]

