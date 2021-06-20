# Build image
FROM golang:alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache bash ca-certificates make git curl build-base

RUN mkdir /admini

WORKDIR /admini

RUN go get -u github.com/pyros2097/go-embed
RUN go get -u github.com/valyala/quicktemplate
RUN go get -u github.com/valyala/quicktemplate/qtc
RUN go get -u golang.org/x/tools/cmd/goimports

ADD ./go.mod          /admini/go.mod
ADD ./go.sum          /admini/go.sum

RUN go mod download

ADD ./app             /admini/app
ADD ./assets          /admini/assets
ADD ./bin             /admini/bin
ADD ./main.go         /admini/main.go
ADD ./Makefile        /admini/Makefile
ADD ./queries         /admini/queries
ADD ./views           /admini/views

RUN go mod download

RUN set -xe && bash -c 'make build-release-ci'

RUN mv build/release /build

# Final image
FROM alpine

RUN apk add --update --no-cache ca-certificates tzdata bash curl

SHELL ["/bin/bash", "-c"]

COPY --from=builder /build/* /usr/local/bin/

EXPOSE 14000
CMD ["admini"]
