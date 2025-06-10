ARG GO_VERSION=1.24.4
ARG TAG=release
ARG VERSION=

FROM golang:${GO_VERSION}-alpine AS build

RUN apk add --update --no-cache git ca-certificates build-base

WORKDIR /app
COPY go.mod go.sum main.go /app/
RUN go mod download

COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/
COPY internal/ /app/internal/

RUN CGO_ENABLED=0 go build -tags "${TAG}" "-ldflags=-X main.version=${VERSION} -s -w -buildid=" -trimpath -o app ./cmd/app

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/app /app/app

ENTRYPOINT [ "/app/app" ]
