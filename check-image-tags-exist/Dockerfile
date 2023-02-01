FROM golang:1.19.5-alpine

RUN apk --no-cache add bash

WORKDIR /app

COPY . /app

# Compile the action
RUN go build -o $GOPATH/bin/tagcheck

RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
