# Builder stage
FROM golang:1.19-alpine3.16 as builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp \
    HOME=/

RUN apk add --no-cache make gcc curl jq

COPY qat_computer /src/
COPY VERSION.txt /src/docs/VERSION.txt
WORKDIR /src/

RUN make build-app

# Runner stage
FROM alpine:3.16 as runner

ENV PROJECT_NAME=qat-computer

LABEL org.opencontainers.image.title="${PROJECT_NAME}" \
      org.opencontainers.image.authors=${GITHUB_ACTOR} \
      org.opencontainers.image.vendor=${GITHUB_REPOSITORY} \
      org.opencontainers.image.source="https://github.com/mickahell/${PROJECT_NAME}" \
      org.opencontainers.image.url="https://github.com/mickahell/${PROJECT_NAME}/tags" \
      org.opencontainers.image.description="Qiskit base docker image." \
      org.opencontainers.image.documentation="https://github.com/mickahell/${PROJECT_NAME}/blob/main/README.md" \
      org.opencontainers.image.os="ubuntu" \
      org.opencontainers.image.version=${GITHUB_REF}

COPY --from=builder /src/qat-computer /usr/local/bin/qat-computer

RUN adduser --uid 1000 --disabled-password --shell /bin/bash user
USER user

VOLUME /etc/qat-computer/conf/
VOLUME /etc/qat-computer/compute/

WORKDIR /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/qat-computer" ]
CMD [ "--help" ]