.PHONY: run_dev_env
run_dev_env:
	docker run -it --rm -v $(shell pwd):/repo -e NIX_USER_CONF_FILES=/repo/nix.conf --workdir /repo tateexon/nix-ubuntu:latest /bin/bash

main:
	go run main.go ./test.gci

test:
	go test -v ./main_test.go

build:
	GOOS=linux GOARCH=amd64 go build -o primeIGT_linux_amd64
	GOOS=windows GOARCH=amd64 go build -o primeIGT_win_amd64.exe

lint:
	golangci-lint --color=always run ./... -v
