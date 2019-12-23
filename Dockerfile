FROM golang:1.13-alpine

COPY . /usr/src/app
WORKDIR /usr/src/app
RUN go build

CMD ./caffeine-notification