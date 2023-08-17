IMAGE ?= quay.io/akaris/crio-repair:latest

.PHONY: build
build:
	go build -tags exclude_graphdriver_btrfs -v -o _output/crio-repair . 

.PHONY: run
run:
	_output/crio-repair

.PHONY: clean
clean:
	rm -f _output/*

.PHONY: container
container:
	podman build -t $(IMAGE) .
