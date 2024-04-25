FROM golang:1.21.6

LABEL authors="Samuel"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV DB_name=people_api
ENV DB_user=db_user
ENV DB_pswd=password
ENV DB_port=3333
ENV DB_host=host.docker.internal

RUN go build -o GO_API

EXPOSE 8088

CMD ["./GO_API"]