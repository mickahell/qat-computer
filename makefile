CONTAINER_NAME?=qatcomputer
PROJECT_NAME?=qat-computer
PYTHON_VERS?=3.9

docker-build:
	docker build . --build-arg FULL=false --file docker/Dockerfile --tag $(CONTAINER_NAME):latest

docker-start:
	docker run --rm \
		--name $(CONTAINER_NAME) \
		-v $(PWD)/qat_computer/resources/conf:/etc/$(PROJECT_NAME)/conf \
		-v $(PWD)/qat_computer/resources/compute:/etc/$(PROJECT_NAME)/compute \
		$(CONTAINER_NAME):latest \
		compute -conf=/etc/$(PROJECT_NAME)/conf/conf_docker.yaml

compose-start:
	docker-compose up && docker-compose rm -f

tox-test:
	tox -epy$(PYTHON_VERS)

tox-fmt:
	black tests
	tox -elint