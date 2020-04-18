FROM golang:1.14.2-alpine3.11 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

RUN set -eux && \
  apk update && \
  apk add --no-cache git && \
  go get github.com/oxequa/realize && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

ENV GO111MODULE on

RUN set -eux && \
  go build -o kimono-app-api

FROM alpine:3.10

WORKDIR /app

COPY --from=build /go/app/kimono-app-api .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/kimono-app-api

CMD ["./kimono-app-api"]
