# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

ADD . /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go env -w GO111MODULE=on

COPY *.go ./

RUN go build -o /docker-gs-ping
EXPOSE 3001

CMD [ "/docker-gs-ping" ]