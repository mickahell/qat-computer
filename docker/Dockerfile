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
FROM ubuntu:22.04 as runner

# For label
ARG GITHUB_ACTOR
ARG GITHUB_REPOSITORY
ARG GITHUB_REF

# If true, install full qiskit libs
ARG FULL

ENV PROJECT_NAME=qat-computer

LABEL org.opencontainers.image.title="${PROJECT_NAME}" \
      org.opencontainers.image.authors="${GITHUB_ACTOR}" \
      org.opencontainers.image.vendor="${GITHUB_REPOSITORY}" \
      org.opencontainers.image.source="https://github.com/${GITHUB_ACTOR}/${PROJECT_NAME}" \
      org.opencontainers.image.url="https://github.com/${GITHUB_ACTOR}/${PROJECT_NAME}/tags" \
      org.opencontainers.image.description="Qiskit base docker image." \
      org.opencontainers.image.documentation="https://github.com/${GITHUB_ACTOR}/${PROJECT_NAME}/blob/main/README.md" \
      org.opencontainers.image.os="ubuntu" \
      org.opencontainers.image.version="${GITHUB_REF}"

COPY --from=builder /src/qat-computer /usr/local/bin/qat-computer
COPY docker/deps.sh /usr/local/bin/deps.sh
COPY docker/requirements.txt /usr/local/bin/requirements.txt
COPY docker/requirements-full.txt /usr/local/bin/requirements-full.txt

RUN apt update -yq
RUN apt dist-upgrade -yq
RUN apt install -yq python3-pip && apt clean -y

RUN pip install --upgrade pip setuptools
RUN pip install -r /usr/local/bin/requirements.txt
RUN /usr/local/bin/deps.sh

VOLUME /etc/qat-computer/conf/
VOLUME /etc/qat-computer/compute/

WORKDIR /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/qat-computer" ]
CMD [ "--help" ]