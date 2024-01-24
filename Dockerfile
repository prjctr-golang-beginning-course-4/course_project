FROM golang:1.21

WORKDIR /app

COPY go .
COPY .env ./.env

RUN go mod download

RUN go build -o main .

CMD ["./main"]