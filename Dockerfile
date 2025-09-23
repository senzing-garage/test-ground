FROM  debian:13.1-slim

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
