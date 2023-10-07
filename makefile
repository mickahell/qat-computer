PROJECT_NAME?=qat-computer
PYTHON_VERS?=3.9

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker run --rm \
		--name $(PROJECT_NAME) \
		-v $(PWD)/qat_computer/resources/conf:/etc/$(PROJECT_NAME)/conf \
		$(PROJECT_NAME):latest \
		-conf=/etc/$(PROJECT_NAME)/conf/conf_docker.yaml -show-config

compose-start:
	docker-compose up && docker-compose rm -f

tox-test:
	tox -epy$(PYTHON_VERS)

tox-fmt:
	black tests
	tox -elint