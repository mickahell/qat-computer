CONTAINER_NAME?=qatcomputer
PROJECT_NAME?=qat-computer

docker-build:
	docker build . --build-arg FULL=false --file docker/Dockerfile --tag $(CONTAINER_NAME):latest

docker-start:
	docker run -d --rm --privileged --cgroupns=host \
		--name $(CONTAINER_NAME) \
		-v /sys/fs/cgroup:/sys/fs/cgroup:rw \
		-v $(PWD)/qat_computer/resources/conf:/etc/$(PROJECT_NAME)/conf \
		-v $(PWD)/qat_computer/resources/compute:/etc/$(PROJECT_NAME)/compute \
		$(CONTAINER_NAME):latest

docker-compute:
	docker exec --tty $(CONTAINER_NAME) \
		$(PROJECT_NAME) compute -conf=/etc/$(PROJECT_NAME)/conf/conf_docker.yaml

docker-inside:
	docker exec -it $(CONTAINER_NAME) bash

compose-start:
	docker compose up -d

tox-test:
	tox -epy$(PYTHON_VERS)

tox-fmt:
	black tests
	tox -elint