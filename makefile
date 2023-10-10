PROJECT_NAME?=qatcomputer
PYTHON_VERS?=3.9

docker-build:
	docker build . --build-arg FULL=false --file docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker run --rm \
		--name $(PROJECT_NAME) \
		-v $(PWD)/qat_computer/resources/conf:/etc/$(PROJECT_NAME)/conf \
		-v $(PWD)/qat_computer/resources/compute:/etc/$(PROJECT_NAME)/compute \
		$(PROJECT_NAME):latest \
		compute -conf=/etc/$(PROJECT_NAME)/conf/conf_docker.yaml

compose-start:
	docker-compose up && docker-compose rm -f

tox-test:
	tox -epy$(PYTHON_VERS)

tox-fmt:
	black tests
	tox -elint