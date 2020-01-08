.PHONY: ensure build clean run run-clean test

ensure:
	dep ensure -v

build:
	env GOOS=linux go build -v

clean:
	[ ! -e ./boilerplate_app ] || rm ./boilerplate_app

run: clean build
	./boilerplate_app

run-clean: ensure clean build
	./boilerplate_app

test:
	go test -v -cover ./...