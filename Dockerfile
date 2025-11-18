FROM  debian:13.2-slim@sha256:9812458f2932ede726468ba07bcb9e51bceb1f0c7f16ee30baa789ccee7cc202

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
