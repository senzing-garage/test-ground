FROM  debian:11.7-slim@sha256:924df86f8aad741a0134b2de7d8e70c5c6863f839caadef62609c1be1340daf5

USER 1001

HEALTHCHECK CMD echo "test ground"

CMD cat /etc/os-release
