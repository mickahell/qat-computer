CONTAINER_NAME?=demo-qatcomputer
BINARY_NAME?=qat-computer

docker-start:
	docker run -d --rm --privileged --cgroupns=host \
		--name $(CONTAINER_NAME) \
		-v /sys/fs/cgroup:/sys/fs/cgroup:rw \
		-v $(PWD)/conf:/etc/$(BINARY_NAME)/conf \
		-v $(PWD)/experiments:/etc/$(BINARY_NAME)/compute \
		mickahell/qatcomputer:latest

docker-compute-qubo:
	docker exec --tty $(CONTAINER_NAME) \
		$(BINARY_NAME) compute -conf=conf_qubo.yaml

docker-compute-qnn:
	docker exec --tty $(CONTAINER_NAME) \
		$(BINARY_NAME) compute -conf=conf_qnn.yaml
