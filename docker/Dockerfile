# Builder stage
FROM golang:1.19-alpine3.16 as builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp \
    HOME=/

RUN apk add --no-cache make gcc curl jq

COPY qat_computer /src/
COPY VERSION.txt /src/docs/
WORKDIR /src/

RUN make build-app

# Runner stage
FROM ubuntu:22.04 as runner

# For label
ARG IMAGE_NAME
ARG GITHUB_ACTOR
ARG GITHUB_REPOSITORY
ARG GITHUB_REF

# If true, install full qiskit libs
ARG FULL

ENV PROJECT_NAME=qatcomputer
ENV CONTAINER_STATUS="true"

LABEL org.opencontainers.image.title="${IMAGE_NAME}" \
      org.opencontainers.image.authors="${GITHUB_ACTOR}" \
      org.opencontainers.image.vendor="${GITHUB_REPOSITORY}" \
      org.opencontainers.image.source="https://github.com/${GITHUB_REPOSITORY}" \
      org.opencontainers.image.url="https://github.com/${GITHUB_REPOSITORY}/tags" \
      org.opencontainers.image.description="Qiskit base docker image." \
      org.opencontainers.image.documentation="https://github.com/${GITHUB_REPOSITORY}/blob/main/README.md" \
      org.opencontainers.image.os="ubuntu" \
      org.opencontainers.image.version="${GITHUB_REF}"

COPY --from=builder /src/qat-computer /usr/local/bin/qat-computer
COPY docker/deps.sh /usr/local/bin/deps.sh
COPY docker/requirements.txt /usr/local/bin/requirements.txt
COPY docker/requirements-full.txt /usr/local/bin/requirements-full.txt

COPY docker/initctl_faker .

RUN apt update -yq
RUN apt dist-upgrade -yq
RUN apt install -yq --no-install-recommends \
    apt-utils build-essential locales lsb-release \
    python3-setuptools python3-pip python3-yaml \
    software-properties-common \
    systemd systemd-cron sudo \
    && apt clean -y

RUN locale-gen en_US.UTF-8

RUN pip install --upgrade pip setuptools
RUN pip install -r /usr/local/bin/requirements.txt
RUN /usr/local/bin/deps.sh

RUN chmod +x initctl_faker && rm -fr /sbin/initctl && ln -s /initctl_faker /sbin/initctl

# Remove unnecessary getty and udev targets that result in high CPU usage when using multiple containers
RUN rm -f /lib/systemd/system/systemd*udev* \
    && rm -f /lib/systemd/system/getty.target

VOLUME ["/sys/fs/cgroup", "/tmp", "/run"]
VOLUME /etc/qat-computer/conf/
VOLUME /etc/qat-computer/compute/

WORKDIR /usr/local/bin/

HEALTHCHECK --interval=10s --timeout=5s \
    CMD if [[ $(systemctl is-system-running) == "running" ]]; then echo 0; else echo 1; fi

CMD ["/lib/systemd/systemd"]