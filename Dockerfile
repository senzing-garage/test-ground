FROM  debian:13.1-slim@sha256:1caf1c703c8f7e15dcf2e7769b35000c764e6f50e4d7401c355fb0248f3ddfdb

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
