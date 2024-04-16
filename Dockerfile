FROM golang:1.21.6

LABEL authors="Samuel"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o GO_API

EXPOSE 8088

CMD ["./GO_API"]