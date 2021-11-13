FROM golang:alpine

RUN apk add --update --no-cache ca-certificates libc6-compat

# main http port
EXPOSE 14000
# marketing port
EXPOSE 14001

ENTRYPOINT ["/admini", "-a", "0.0.0.0"]

COPY admini /
