GO111MODULE?=on

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

pull-app: setting-prepare
	go run . pull -image ghcr.io/mickahell/qatcomputer-dev -tag latest

start-app: setting-prepare
	go run . init -compute-volume ./../qat_computer/resources/compute -conf-volume ./../qat_computer/resources/conf

compute-app: setting-prepare
	go run . qatcomputer -computer-args "compute -conf /etc/$(PROJECT_NAME)/conf/conf_docker.yaml"

stop-app: setting-prepare
	go run . stop -remove

build-app: setting-prepare
	go build -mod=vendor

install-app: build-app
	go install