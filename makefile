PROJECT_NAME?=qat-computer

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker run --rm \
		--name $(PROJECT_NAME) \
		-v $(PWD)/qat_computer/conf:/etc/$(PROJECT_NAME)/conf \
		$(PROJECT_NAME):latest \
		-conf=/etc/$(PROJECT_NAME)/conf/conf_local.yaml