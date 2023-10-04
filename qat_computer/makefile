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

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build -mod=vendor

install-app: build-app
	go install