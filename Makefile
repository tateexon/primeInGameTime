.PHONY: run_dev_env
run_dev_env:
	docker run -it --rm -v $(shell pwd):/repo -e NIX_USER_CONF_FILES=/repo/nix.conf --workdir /repo tateexon/nix-ubuntu:latest /bin/bash
