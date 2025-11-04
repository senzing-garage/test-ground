FROM  debian:13.1-slim@sha256:a347fd7510ee31a84387619a492ad6c8eb0af2f2682b916ff3e643eb076f925a

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
