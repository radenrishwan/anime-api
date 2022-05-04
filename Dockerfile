FROM golang:1.18.0

WORKDIR /app

COPY go.mod go.mod ./

RUN apk add chromium

RUN go mod download

COPY . .

RUN apk add git

RUN go build -o main .

EXPOSE 3000

CMD ./main

