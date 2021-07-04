FROM scratch
ENTRYPOINT ["/admini", "-a", "0.0.0.0"]
EXPOSE 14000
COPY admini /
