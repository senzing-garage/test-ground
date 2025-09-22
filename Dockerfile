ARG IMAGE_SENZINGAPI_RUNTIME=senzing/senzingapi-runtime:3.12.6

FROM ${IMAGE_SENZINGAPI_RUNTIME} AS senzingapi-runtime

# FROM  debian:13.1-slim

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
