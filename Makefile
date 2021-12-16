default: build

.PHONY: build test

build:
	go build

test:
	go test ./pkg/netkit ./driver/uml ./driver/podman ./util/topsort
