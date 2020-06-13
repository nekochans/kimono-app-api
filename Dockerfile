FROM golang:1.14.4-alpine3.11 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

ARG GOLANGCI_LINT_VERSION=v1.27.0

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl && \
  curl -fLo /go/bin/air https://git.io/linux_air && \
  chmod +x /go/bin/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}

ENV GO111MODULE on

RUN set -eux && \
  go build -o kimono-app-api

ENV CGO_ENABLED 0

FROM alpine:3.10

WORKDIR /app

COPY --from=build /go/app/kimono-app-api .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/kimono-app-api

CMD ["./kimono-app-api"]
