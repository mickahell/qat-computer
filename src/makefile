GO111MODULE?=on
PROJECT_NAME?=qatcomputer
PROJECT_PATH?=QatComputer

setting-prepare:
	go mod tidy
	go mod vendor

install-devtools: setting-prepare
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	go install github.com/segmentio/golines@latest

run-test:
	go clean -testcache
	go test ./... -covermode=count -coverprofile ./coverage.out

run-fmt:
	go fmt
	$(HOME)/go/bin/yamlfmt conf/
	$(HOME)/go/bin/golines . -w

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build -mod=vendor

install-app: build-app
	go install

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker run --rm \
		--name $(PROJECT_NAME) \
		-v $(PWD)/conf:/etc/$(PROJECT_PATH)/conf \
		-e PASS_DB=${PASS_DB} -e TOKEN_API_TEST=${TOKEN_API_TEST} \
		-p 8080:8080 \
		$(PROJECT_NAME):latest \
		-conf=/etc/$(PROJECT_PATH)/conf/conf_docker.yaml