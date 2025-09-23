ARG IMAGE_SENZINGAPI_RUNTIME=senzing/senzingapi-runtime:3.12.8@sha256:3663a1971e564af4d12ecdb0c90a4f46418b77dc229ec6c9f692efc59d1c67ae

FROM ${IMAGE_SENZINGAPI_RUNTIME} AS senzingapi-runtime

# FROM  debian:13.1-slim

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
