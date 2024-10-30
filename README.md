# Qat Computer

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/mickahell/qat-computer)](https://github.com/mickahell/qat-computer/releases)

[![Demo tests](https://github.com/mickahell/qat-computer/actions/workflows/demo.yml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/demo.yml)
[![Python tests](https://github.com/mickahell/qat-computer/actions/workflows/python.yaml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/python.yaml)
[![QatComputer tests](https://github.com/mickahell/qat-computer/actions/workflows/qat-computer.yaml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/qat-computer.yaml)
[![QatCMD tests](https://github.com/mickahell/qat-computer/actions/workflows/qat-cmd.yaml/badge.svg)](https://github.com/mickahell/qat-computer/actions/workflows/qat-cmd.yaml)

[![Docker Pulls](https://img.shields.io/docker/pulls/mickahell/qatcomputer?label=QatComputer&style=for-the-badge)](https://hub.docker.com/r/mickahell/qatcomputer)
[![Docker Pulls Full](https://img.shields.io/docker/pulls/mickahell/qatcomputer-full?label=QatComputer%20Full&style=for-the-badge)](https://hub.docker.com/r/mickahell/qatcomputer-full)

[![Qiskit](https://img.shields.io/badge/Qiskit-%E2%89%A5%201.0.0-6133BD)](https://github.com/Qiskit/qiskit)
[![License](https://img.shields.io/github/license/qiskit-community/quantum-prototype-template?label=License)](https://github.com/IceKhan13/purplecaffeine/blob/main/LICENSE)

The project is aimed to create a Qiskit base docker image to use as base for any quantum application using Qiskit.
The image also contains the tool `QatComputer` who allows you to run a python project directly inside the container as a `run as you go`.

Full documentation available at [QatComputer docs](https://mickahell.github.io/qat-computer/).

## Goals

This project has 2 purposes :
- Create a based docker image pre-installed with an updated version of Qiskit. In order to help with integration and deployement of Quantum apps.
- Create a virtual environment to run and test your own project locally, serverlessly without having to configure anything or to alter your local configuration. AS everything is inside a docker env, you create, run, destroy and retry again with a clean new container again and again.

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
docker pull mickahell/qatcomputer
# Full
docker pull mickahell/qatcomputer-full
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
	-v my_conf_folder:/etc/qat-computer/conf \
	-v my_super_project:/etc/qat-computer/compute \
	mickahell/qatcomputer:latest
```

In order to the tool works correctly you need :
- (Optional) Be in priviledge mode and host network mode
- (Optional) Link your cgroup
- (Optional) Link your project conf
- Link your project to the compute folder inside the image

And run the tool :

```bash
docker exec --tty qatcomputer \
	qat-computer compute -conf=conf.yaml
```
