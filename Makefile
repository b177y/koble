default: build

GO ?= go
BUILDTAGS ?= exclude_graphdriver_btrfs btrfs_noversion exclude_graphdriver_devicemapper containers_image_openpgp
BUILDFLAGS := -mod=vendor $(BUILDFLAGS)

.PHONY: build test

build:
	$(GO) build $(BUILDFLAGS) -tags "$(BUILDTAGS)" -o bin/netkit

test:
	$(GO) test $(BUILDFLAGS) -tags "$(BUILDTAGS)" ./pkg/netkit ./driver/uml ./driver/podman ./util/topsort

.PHONY: vendor
vendor:
	GO111MODULE=on $(GO) mod tidy
	GO111MODULE=on $(GO) mod vendor
	GO111MODULE=on $(GO) mod verify
	rm vendor/github.com/containers/storage/pkg/unshare/unshare_cgo.go
