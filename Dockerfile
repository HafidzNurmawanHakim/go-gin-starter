FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main ./cmd

FROM golang:1.23

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .env

EXPOSE 8080

CMD ["/app/main"]
