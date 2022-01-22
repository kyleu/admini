# Content managed by Project Forge, see [projectforge.md] for details.
FROM golang:alpine

LABEL "org.opencontainers.image.authors"="Kyle U"
LABEL "org.opencontainers.image.source"="https://github.com/kyleu/admini"
LABEL "org.opencontainers.image.vendor"="kyleu"
LABEL "org.opencontainers.image.title"="Admini"
LABEL "org.opencontainers.image.description"="Use Admini to explore and manage your data as fast and easily as possible"

RUN apk add --update --no-cache ca-certificates libc6-compat

# main http port
EXPOSE 14000
# marketing port
EXPOSE 14001

ENTRYPOINT ["/admini", "-a", "0.0.0.0"]

COPY admini /
