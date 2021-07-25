FROM golang:1.16-alpine3.14

WORKDIR /home/basic

RUN apk update && apk upgrade \
  && apk --no-cache add curl \
  && apk add --no-cache --upgrade bash

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./basic/go.mod .
COPY ./basic/go.sum .
COPY ./basic/public ./public

RUN go mod download

CMD [ "air" ]