FROM golang:1.18.0-alpine

WORKDIR /app

COPY go.mod go.mod ./

RUN go mod download

COPY . .

RUN apk add git

RUN go build -o main .

EXPOSE 3000

CMD ./main

