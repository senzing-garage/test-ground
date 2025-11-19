FROM  debian:13.2-slim@sha256:18764e98673c3baf1a6f8d960b5b5a1ec69092049522abac4e24a7726425b016

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
