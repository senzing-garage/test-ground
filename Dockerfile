FROM  debian:13.1-slim@sha256:66b37a5078a77098bfc80175fb5eb881a3196809242fd295b25502854e12cbec

USER 1001

HEALTHCHECK CMD echo "test ground"

ENTRYPOINT ["cat","/etc/os-release"]
