.PHONY: run_dev_env
run_dev_env:
	docker run -it --rm -v $(shell pwd):/repo -e NIX_USER_CONF_FILES=/repo/nix.conf --workdir /repo tateexon/nix-ubuntu:latest /bin/bash

main:
	go run main.go ./test.gci

test:
	go test test.go

build:
	GOOS=linux GOARCH=amd64 go build -o primeIGT_linux_amd64
