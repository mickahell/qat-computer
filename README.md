# Qat Computer

[![Python tests](https://github.com/mickahell/qat-computer/actions/workflows/python.yaml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/python.yaml)
[![QatComputer tests](https://github.com/mickahell/qat-computer/actions/workflows/qat-computer.yaml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/qat-computer.yaml)
[![QatCMD tests](https://github.com/mickahell/qat-computer/actions/workflows/qat-cmd.yaml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/qat-cmd.yaml)

[![Docker Pulls](https://img.shields.io/docker/pulls/mickahell/qatcomputer?label=QatComputer&style=for-the-badge)](https://hub.docker.com/r/mickahell/qatcomputer)
[![Docker Pulls](https://img.shields.io/docker/pulls/mickahell/qatcomputer-full?label=QatComputer%20Full&style=for-the-badge)](https://hub.docker.com/r/mickahell/qatcomputer-full)

<!--
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/mickahell/qat-computer)](https://github.com/mickahell/qat-computer/releases)
[![Docker Pulls](https://img.shields.io/docker/pulls/mickahell/qat-computer?label=QatComputer&style=for-the-badge)](https://hub.docker.com/r/mickahell/qat-computer)
-->

The project is aimed to create a Qiskit base docker image to use as base for any quantum application using Qiskit.
The image also contains the tool `QatComputer` who allows you to run a python project directly inside the container as a `run as you go`.

Full documentation available at [QatComputer docs](https://mickahell.github.io/qatcomputer/).

## Prerequisites

- <details><summary>Linux</summary>
  <pre>apt-get install docker-ce docker-ce-cli containerd.io</pre>
</details>

- <details><summary>Mac / Windows</summary>
  https://www.docker.com/products/docker-desktop
</details>

## Details

- The image is base on Ubuntu [The Jammy Jellyfish](https://doc.ubuntu-fr.org/jammy)
- Python version install : [3.10](https://www.python.org/downloads/release/python-3104/)

## Getting started

The image is available in DockerHub and can be clone with :

```bash
docker clone ...
```

Most of the usefull commands are available in the [Docker Compose file](./docker-compose.yml) and in the [makefile](./makefile).

### Build

if you want to build your own image :

```bash
docker build . --build-arg FULL=False --file docker/Dockerfile --tag qatcomputer:latest
```

### Compute

Parameters for the tool can be given as docker extra parameters. A configuration file is also available with every parameters :

```yaml
# log
## 0 : Informational
## 1 : Warning
## 2 : Error --> Always
## 3 : Critical --> Always
loglevel: 0   # Optionnal
debian_packages:   # Optionnal
  - "jq"
python_version: "python3.10"   # Optionnal
compute_path: "simple_python_repository"
filename_to_execute: "main.py"
requirements_file: "requirements.txt"   # Optionnal
```

In order to start the image and make the internal tool available :

```bash
docker run -d --rm --privileged --cgroupns=host \
	--name qatcomputer \
	-v /sys/fs/cgroup:/sys/fs/cgroup:rw \
	-v $(PWD)/qat_computer/resources/conf:/etc/qat-computer/conf \
	-v $(PWD)/qat_computer/resources/compute:/etc/qat-computer/compute \
	qatcomputer:latest
```

And run the tool :

```bash
docker exec --tty qatcomputer \
	qat-computer compute -conf=conf_docker.yaml
```
