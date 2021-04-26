FROM golang:1.16-alpine as build

WORKDIR /go/src/app

ADD . .

RUN apk update \
    && apk add git \
    && GO111MODULE=off go get github.com/oxequa/realize \
    && go get -v ./... \
    && go build -o ./binary/app ./cmd

FROM alpine

WORKDIR /app

COPY --from=build /go/src/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]