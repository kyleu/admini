FROM golang:alpine

RUN apk add --update --no-cache ca-certificates

ENTRYPOINT ["/admini", "-a", "0.0.0.0"]
EXPOSE 14000
COPY admini /
